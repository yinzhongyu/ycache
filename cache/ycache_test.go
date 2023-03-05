package cache

import (
	"log"
	"testing"
)

func TestYCache_Put(t *testing.T) {
	log.Println("test the YCache")
	k1, k2, k3 := "key1", "key2", "key3"
	v1, v2, v3 := []byte("value1"),[]byte("value2"), []byte("value3")

 	cache := NewYcache(3)

	cache.Put(k1, v1)
	cache.Put(k2, v2)
	cache.Put(k3, v3)

	if v, ok := cache.Get("key1"); !ok || string(v) != string(v1)  {
		t.Fatalf("k1 put cache  failed")
	}
	if v, ok := cache.Get("key2"); !ok || string(v) !=  string(v2) {
		t.Fatalf("k2 put cache  failed")
	}
	if v, ok := cache.Get("key3"); !ok || string(v) !=  string(v3) {
		t.Fatalf("k3 put cache  failed")
	}

	k4, v4 := "key4", []byte("value4")
	cache.Put(k4, v4)

	if _, ok := cache.Get("key1"); ok {
		t.Fatalf("cache hit key1 failed")
	}

	if v, ok := cache.Get("key4"); !ok || string(v) !=  string(v4) {
		t.Fatalf("k4 put cache  failed")
	}


}
