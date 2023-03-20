package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"ycache/consistenthash"
	"ycache/ycachepb"
)

const defaultReplicas = 50
const basePath = "/ycache"
// PeerPicker is the interface that must be implemented to locate
// the peer that owns a specific key.
type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// PeerGetter is the interface that must be implemented by a peer.
type PeerGetter interface {
	Get(in *ycachepb.Request, out *ycachepb.Response) error
}


type httpClient struct {
	basePath string
	mu          sync.Mutex // guards peers and httpGetters
	peers       *consistenthash.Map
	httpGetters map[string]*httpGetter // keyed by e.g. "http://10.0.0.2:8008"
}

type httpGetter struct {
	baseURL string
}

func NewHttpClient()*httpClient{
	return &httpClient{basePath: basePath,peers: consistenthash.NewMap(defaultReplicas,func(key []byte) uint32 {
		i, _ := strconv.Atoi(string(key))
		return uint32(i)
	}),httpGetters: make(map[string]*httpGetter)}
}
func (h *httpGetter) Get(in *ycachepb.Request, out *ycachepb.Response) error {
	u := fmt.Sprintf(
		"%v%v/%v",
		h.baseURL,
		url.QueryEscape(in.GetGroup()),
		url.QueryEscape(in.GetKey()),
	)
	res, err := http.Get(u)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return  fmt.Errorf("server returned: %v", res.Status)
	}

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("reading response body: %v", err)
	}

	if err = proto.Unmarshal(bytes, out); err != nil {
		return fmt.Errorf("decoding response body: %v", err)
	}

	return  nil
}

var _ PeerGetter = (*httpGetter)(nil)


// Set updates the pool's list of peers.
func (p *httpClient) Set(peers []string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.peers = consistenthash.NewMap(defaultReplicas, func(key []byte) uint32 {
		i, _ := strconv.Atoi(string(key))
		return uint32(i)
	})
	log.Println(111,p.peers)
	p.peers.Add(peers)
	p.httpGetters = make(map[string]*httpGetter, len(peers))
	for _, peer := range peers {
		p.httpGetters[peer] = &httpGetter{baseURL: peer + p.basePath}
	}
}

// PickPeer picks a peer according to key
func (p *httpClient) PickPeer(key string) (PeerGetter, bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if peer := p.peers.Get(key); peer != "" {
		log.Printf("Pick peer %s", peer)
		return p.httpGetters[peer], true
	}
	return nil, false
}

var _ PeerPicker = (*httpClient)(nil)

