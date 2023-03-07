package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

type Hash func(data []byte)uint32


type Map struct {
	num int //虚拟节点的倍数
	keys []int  //环（有序）
 	names map[int]string  //数字与名字的映射
	hash Hash
}

//
// NewMap
//  @Summary 摘要
//  @Description: 新建一个一致性哈希环
//  @input: 虚拟节点倍数及自定义的哈希函数
//  @output: 一致性哈希环
func NewMap(num int,hash Hash)*Map{
	m:= &Map{num: num,keys: []int{},names: make(map[int]string)}
	m.hash = hash
	if m.hash == nil {
		hash = crc32.ChecksumIEEE
	}














	//if hash == nil {
	//		hash = crc32.ChecksumIEEE
	//}else{
	//	m.hash = hash
	//}

	return m
}


//
// Add
//  @Summary 摘要
//  @Description: 根据节点名字，生成虚拟节点，并存入环上数字和名字的映射
//  @input: names 节点名字
//  @output:
func (m *Map) Add(names []string){
	for i := 0; i < len(names); i++ {
		for j := 0; j < m.num; j++ {
			tempName := strconv.Itoa(j) + string(names[i])
			tempNum:=m.hash([]byte(tempName))
			m.keys = append(m.keys,int(tempNum))
			m.names[int(tempNum)] = names[i]
		}
	}
	sort.Ints(m.keys)
}

//
// Get
//  @Summary 摘要
//  @Description: 根据传入的key值，确定需要访问哪个节点
//  @input: key
//  @output: 节点
func (m *Map) Get(key string)string{
	hashNum := int(m.hash([]byte(key)))
	for i := 0; i < len(m.keys); i++ {
		if hashNum <= m.keys[i]{
			return m.names[m.keys[i]]
		}
	}

	return m.names[m.keys[0]]
}
