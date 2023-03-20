package singleflight

import "sync"

type Group struct {
	mu sync.Mutex
	m map[string]*Call
}

// 存储每次调用的函数参数
type Call struct {
	wg sync.WaitGroup
	val interface{}
	err error
}

func (g *Group) Do(key string,fn func()(interface{},error))(interface{},error) {

	g.mu.Lock()
	//懒加载
	if g.m == nil{
		g.m = make(map[string]*Call)
	}

	if v,ok:=g.m[key];ok{
		g.mu.Unlock()
		v.wg.Wait()
		return v.val,v.err
	}else {
		c := &Call{}

		c.wg.Add(1)
		g.m[key] = c
		g.mu.Unlock()

		c.val,c.err = fn()
		c.wg.Done()

		g.mu.Lock()
		delete(g.m,key)
		g.mu.Unlock()

		return c.val,c.err
	}

}

