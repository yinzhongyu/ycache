package singleflight

import (
	"log"
	"sync"
	"testing"
	"time"
)



func TestGroup_Do(t *testing.T) {
	var wg sync.WaitGroup
	g := &Group{}
	var num int
	//模仿Db函数
	getDb := func(i int,key string)[]byte {
		log.Println(i,"search the Db for",key)
		num++
		time.Sleep(1*time.Second)
		return []byte(key)
	}

	key := "key"
	for i := 0; i < 10; i++ {
	go func(i int) {
		wg.Add(1)
		v,_:=g.Do(key, func() (interface{}, error) {
			return getDb(i,key),nil
		})
		log.Println(i,"get the data :",v)
		wg.Done()
	}(i)
	}
	time.Sleep(2*time.Second)

	for i := 10; i < 20; i++ {
		go func(i int) {
			wg.Add(1)
			v,_:=g.Do(key, func() (interface{}, error) {
				return getDb(i,key),nil
			})
			log.Println(i,"get the data :",v)
			wg.Done()
		}(i)
	}
	wg.Wait()

	if num > 2 {
		log.Fatal("wrong!there are two many getDb!")
	}
}