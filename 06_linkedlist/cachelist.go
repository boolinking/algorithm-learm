package main


type LRULinkedList struct {
	head *Node
	len uint //链表的长度
	cap uint
}

func NewLRULinkedList(cap uint) *LRULinkedList {
	if cap <= 10 {
		cap = 10 //缓存默认容量
	}
	return &LRULinkedList{nil,0,cap}
}

func (cache *LRULinkedList) Add(data interface{})  {
	//如果已存在，删除节点，放入头部
	ok := cache.Find(data)
	if ok {
		cache.Delete(data)
		cache.InsertHead(data)
	} else if cache.len == cache.cap {//缓存已满,删除尾部节点
		cache.DeleteTail()
		cache.InsertHead(data)
	} else {
		cache.InsertHead(data)
	}
}

func (cache *LRULinkedList) getCache() (interface{}, bool) {
	if cache.head == nil {
		return nil, false
	}
	return cache.head.data, true
}

func (cache *LRULinkedList) InsertHead(data interface{})  {
	n := &Node{data,nil}
	if cache.head == nil {
		cache.head = n
	} else {
		n.next = cache.head
		cache.head = n
	}
}

func (cache *LRULinkedList) Find(data interface{}) bool {
	if cache.head == nil {
		return false
	}
	p := cache.head
	for p.data != data  {
		p = p.next
		if p == nil{
			return false
		}
	}
	return true
}

//删除
func (cache *LRULinkedList) Delete(data interface{})  {
	p := cache.head
	for p != nil && p.next.data != data {
		p = p.next
	}

	p.next  = p.next.next
}

//删除为尾节点
func (cache *LRULinkedList) DeleteTail()  {
	p := cache.head
	for p != nil && p.next.next != nil {
		p = p.next
	}
	p.next  = nil
}
