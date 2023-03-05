package main

import (
	"fmt"
	"log"
	"sync"
	"ycache/cache"
)



// A Getter loads data for a key.
type Getter interface {
	Get(key string) ([]byte, error)
}

// A GetterFunc implements Getter with a function.
type GetterFunc func(key string) ([]byte, error)

// Get implements Getter interface function
func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}





// A Group is a cache namespace and associated data loaded spread over
type Group struct {
	name      string
	getter    Getter
	mainCache *cache.YCache
}

var (
	mu     sync.RWMutex
	groups = make(map[string]*Group)
)

// NewGroup create a new instance of Group
func NewGroup(name string, cap int, getter Getter) *Group {
	if getter == nil {
		panic("nil Getter")
	}
	mu.Lock()
	defer mu.Unlock()
	g := &Group{
		name:      name,
		getter:    getter,
		mainCache: cache.NewYcache(cap),
	}

	groups[name] = g
	return g
}

// GetGroup returns the named group previously created with NewGroup, or
// nil if there's no such group.
func GetGroup(name string) *Group {
	mu.RLock()
	g := groups[name]
	mu.RUnlock()
	return g
}



// Get value for a key from cache
func (g *Group) Get(key string) ([]byte, error) {
	if key == "" {
		return []byte{}, fmt.Errorf("key is required")
	}

	if v, ok := g.mainCache.Get(key); ok {
		log.Println("[GeeCache] hit")
		return v, nil
	}else{
		log.Println("[GeeCache]" ,key,"not hit")
		return g.load(key)
	}

}

func (g *Group) load(key string) (value []byte, err error) {
	return g.getLocally(key)
}

func (g *Group) getLocally(key string) ([]byte, error) {
	bytes, err := g.getter.Get(key)
	if err != nil {
		return []byte{},err
	}
	//存入本地
	g.mainCache.Put(key, bytes)
	return bytes, nil
}

func (g *Group) populateCache(key string, value []byte) {
	g.mainCache.Put(key, value)
}

