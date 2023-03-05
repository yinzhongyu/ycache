package main

import (
	"fmt"
	"log"
	"testing"
)
//
//func TestGetGroup(t *testing.T) {
//
//	myGroup := NewGroup("yzy",5,GetterFunc(func(key string) ([]byte, error) {
//		return []byte(key),nil
//	}))
//	myGroup.Get("hello")
//	myGroup.Get("hello")
//
//
//}


func TestGroup_Get(t *testing.T) {
	var db = map[string]string{
		"Tom":  "630",
		"Jack": "589",
		"Sam":  "567",
	}

	loadCounts := make(map[string]int)

	gee := NewGroup("scores", 5, GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				loadCounts[key] += 1
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))


	for k, _ := range db {
		if _, err := gee.Get(k); err != nil {
			t.Fatal("failed to get value of Tom")
		}else{
		} // load from callback function

		if _, err := gee.Get(k); err != nil || loadCounts[k] > 1 {
			t.Fatalf("cache %s miss", k)
		} // cache hit
	}

	if view, err := gee.Get("unknown"); err == nil {
		t.Fatalf("the value of unknow should be empty, but %s got", view)
	}

}


