package cache

import "testing"

func TestYCachet(t *testing.T) {
	k1, k2, k3 := "key1", "key2", "key3"
	v1, v2, v3 := "value1", "value2", "value3"

	cache := NewYCache(3)
	cache.Put(k1, v1)
	cache.Put(k2, v2)
	cache.Put(k3, v3)

	if v1 != cache.Get(k1) || v2 != cache.Get(k2) || v3 != cache.Get(k3) {
		t.Fatal("put wrong")
	}
	k4, v4 := "key4", "value4"
	cache.Put(k4, v4)
	if cache.Get(k1) != "" {
		t.Fatal("k1 is still on cache")
	}
	if cache.Get(k4) != v4 {
		t.Fatal("k4 put wrong ")
	}
}
