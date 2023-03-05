package cache

type cache struct {
	data   map[string]*ListNode
	length int
	cap    int
	head   *ListNode
	Tail   *ListNode
}

type ListNode struct {
	Key  string
	Val  []byte
	Next *ListNode
	Pre  *ListNode
}

//
// Newcache
//  @Summary 摘要
//  @Description: 生成一个Cache
//  @input: cap(容量)
//  @output:*cache
func Newcache(cap int) *cache {
	head := &ListNode{}
	tail := &ListNode{}

	head.Next = tail
	tail.Next = head

	return &cache{head: head, Tail: tail, cap: cap, length: 0, data: make(map[string]*ListNode)}
}

//
// Get
//  @Summary 摘要
//  @Description: 从缓存中获取value，如果key不存在，返回""
//  @input: key值 (string)
//  @output:value值([]byte)
func (c *cache) Get(key string) ([]byte,bool) {
	if v, ok := c.data[key]; !ok {
		return nil,false
	} else {
		c.setNodeToHead(v)
		return v.Val,true
	}
}

//
// Put
//  @Summary 摘要
//  @Description: 向缓存中存入kv值，如果未满直接插入，如果已满，需移除尾部节点再插入
//  @input:key值、value值(string,[]byte)
//  @output:
func (c *cache) Put(key string, val []byte) {
	if v, ok := c.data[key]; ok {
		v.Val = val
		return
	}

	node := &ListNode{Val: val, Key: key}
	if c.length < c.cap {
		c.insertNodeToHead(node)
	} else {
		c.deleteTailNode()
		c.insertNodeToHead(node)
	}
}

//
// insertNodeToHead
//  @Summary 摘要
//  @Description: 向头部插入节点
//  @input:node
//  @output:
func (c *cache) insertNodeToHead(node *ListNode) {
	c.data[node.Key] = node

	node.Pre = c.head
	node.Next = c.head.Next

	c.head.Next = node
	node.Next.Pre = node

	c.length++
}

//
// setNodeToHead
//  @Summary 摘要
//  @Description: 设置中间节点到头部
//  @input: node
//  @output:
func (c *cache) setNodeToHead(node *ListNode) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre

	c.insertNodeToHead(node)

	c.length--
}

//
// deleteTailNode
//  @Summary 摘要
//  @Description: 删除尾部节点
//  @input:
//  @output:
func (c *cache) deleteTailNode() {

	delete(c.data, c.Tail.Pre.Key)

	c.Tail.Pre.Pre.Next = c.Tail
	c.Tail.Pre = c.Tail.Pre.Pre

}
