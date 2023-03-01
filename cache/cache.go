package cache

type YCache struct {
	data   map[string]*ListNode
	length int
	cap    int
	head   *ListNode
	Tail   *ListNode
}

type ListNode struct {
	Key  string
	Val  string
	Next *ListNode
	Pre  *ListNode
}

//
// NewYCache
//  @Summary 摘要
//  @Description: 生成一个Cache
//  @input: cap(容量)
//  @output:*cache
func NewYCache(cap int) *YCache {
	head := &ListNode{}
	tail := &ListNode{}

	head.Next = tail
	tail.Next = head

	return &YCache{head: head, Tail: tail, cap: cap, length: 0, data: make(map[string]*ListNode)}
}

//
// Get
//  @Summary 摘要
//  @Description: 从缓存中获取value，如果key不存在，返回""
//  @input: key值 (string)
//  @output:value值(string)
func (c *YCache) Get(key string) string {
	if v, ok := c.data[key]; !ok {
		return ""
	} else {
		c.setNodeToHead(v)
		return v.Val
	}
}

//
// Put
//  @Summary 摘要
//  @Description: 向缓存中存入kv值，如果未满直接插入，如果已满，需移除尾部节点再插入
//  @input:key值、value值(string,string)
//  @output:
func (c *YCache) Put(key string, val string) {
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
func (c *YCache) insertNodeToHead(node *ListNode) {
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
func (c *YCache) setNodeToHead(node *ListNode) {
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
func (c *YCache) deleteTailNode() {

	delete(c.data, c.Tail.Pre.Key)

	c.Tail.Pre.Pre.Next = c.Tail
	c.Tail.Pre = c.Tail.Pre.Pre

}
