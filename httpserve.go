package main

import (
	"google.golang.org/protobuf/proto"
	"net/http"
	"strings"
	"ycache/ycachepb"
)



func  httpserver(w http.ResponseWriter, r *http.Request) {

	parts := strings.SplitN(r.URL.Path[len("/ycache")+1:], "/", 2)

	if len(parts) != 2 {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	groupName := parts[0]
	key := parts[1]

	group := GetGroup(groupName)
	if group == nil {
		http.Error(w, "no such group: "+groupName, http.StatusNotFound)
		return
	}

	view, err := group.Get(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Write the value to the response body as a proto message.
	body, err := proto.Marshal(&ycachepb.Response{Value: view})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(body)
	//w.Header().Set("Content-Type", "application/octet-stream")


}
//
//var db = map[string]string{
//	"Tom":  "630",
//	"Jack": "589",
//	"Sam":  "567",
//}
//
//func main() {
//	NewGroup("scores", 5, GetterFunc(
//		func(key string) ([]byte, error) {
//			log.Println("[SlowDB] search key", key)
//			if v, ok := db[key]; ok {
//				return []byte(v), nil
//			}
//			return nil, fmt.Errorf("%s not exist", key)
//		}))
//
//
//	_ = http.ListenAndServe("localhost:8000",http.HandlerFunc(httpserver))
//}
//
//

