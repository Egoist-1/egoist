package lru

import "container/list"

//LRU(Least Recently Used)最近最少访问
/*
	map 存储 键 和 值 的映射
	双向链表实现的队列将所有的值放到链表中 当访问某个值时,将其移动到队尾
*/

type Cache struct {
	maxBytes  int64                         //最大内存数量
	nBytes    int64                         //当前使用内存
	ll        *list.List                    //
	cache     map[string]*list.Element      //存储了节点的
	onEvicted func(key string, value Value) //删除节点的回调函数
}

// entry 双向链表节点的数据
type entry struct {
	key   string
	value Value
}

// Value
type Value interface {
	len() int
}

func New(maxBytes int64,onEvicted func(string,Value)) *Cache{
	return &Cache{
		maxBytes: maxBytes,
		ll:	list.New(),
		cache: make(map[string]*list.Element),
		onEvicted: onEvicted,
	}
}


// Get 查找 从字典找到对应的双向链表节点,然后移动到队尾
func (c *Cache)Get(key string)(value Value,ok bool){
	if ele,ok := c.cache[key];ok{
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value,true 
	}
	return
}

// Add 新增/修改
/* 
	
 */
func (c *Cache)Add(key string,value Value){
	
}

// RemoveOldest 删除(移除最近最少访问的节点(队首))
