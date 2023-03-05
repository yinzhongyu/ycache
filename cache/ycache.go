package cache

import "sync"

type YCache struct {
	cache *cache
	mu sync.Mutex
	Cap int
}

func NewYcache(cap int)*YCache{
	return &YCache{cache: Newcache(cap),Cap: cap}
}
//
// Put
//  @Summary 摘要
//  @Description: 并发存入cache
//  @input:
//  @output:
func (c *YCache) Put(key string,val []byte)  {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache.Put(key,val)
}

//
// Get
//  @Summary 摘要
//  @Description: 并发取cache
//  @input:
//  @output:
func (c *YCache) Get(key string)([]byte,bool){
	c.mu.Lock()
	defer c.mu.Unlock()
  	return c.cache.Get(key)
}