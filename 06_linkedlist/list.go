package main

import (
	"fmt"
)

type Node struct {
	data interface{}
	next *Node
}

type List struct {
	head *Node
}

func NewList(data interface{}) *List {
	head := &Node{data:data,next:nil}
	return &List{head}
}

func (list *List) FindByValue(value interface{}) *Node {
	p := list.head
	for p != nil && p.data != value {
		p = p.next
	}
	return p
}

func (list *List) FindByIndex(index uint) *Node {
	p := list.head
	var i uint = 1
	for ; i != index; i++ {
		p = p.next
	}
	return p
}

//表头插入
func (list *List) InsertToHead(data interface{})  {
	n := &Node{data,nil}
	if list.head == nil {
		list.head = n
	} else {
		n.next = list.head
		list.head = n
	}
}

//在每个节点之前插入
func (list *List) InsertBefore(p *Node , data interface{})  {
	if p == nil {
		return
	}

	if list.head == p {
		list.InsertToHead(data)
		return
	}
	newNode := &Node{data:data,next:nil}
	pre := list.head
	for pre.next != p  {
		pre = pre.next
	}
	newNode.next , pre.next = p , newNode

}

//在某个节点之后插入
func (list *List) InsertAfter( p *Node , data interface{})  {
	if p == nil {
		return
	}
	n := &Node{data:data}
	n.next , p.next = p.next, n
}

//在链表的尾部插入
func (list *List) InsertTail(data interface{})  {
	n := &Node{data,nil}
	p := list.head
	if p == nil {
		list.head = n
		return
	}
	for p.next != nil  {
		p = p.next
	}
	p.next = n
}

//删除指定节点
func (list *List) Delete(data interface{})  {
	p := list.head
	for p != nil && p.next.data != data {
		p = p.next
	}

	p.next  = p.next.next
}
//单链表反转
func (list *List) ReverseList() *List  {
	if list.head == nil {
		return nil
	}
	p := list.head
	if p.next == nil {
		return list
	}
	newList:= NewList(nil)
	for p != nil {
		newList.InsertToHead(p.data)
		p = p.next
	}
	return newList
}

//打印链表
func (list *List) Print() {
	cur := list.head
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.data)
		cur = cur.next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}

func main()  {

	newList := NewList("a")
	newList.InsertToHead("b")
	newList.InsertToHead("c")
	newList.Print()

	newList.InsertTail("b")
	newList.Print()
	n := newList.FindByValue("a")
	newList.InsertBefore(n,"A")
	newList.InsertAfter(n,"B")

	newList.Print()

	l := newList.ReverseList()

	l.Print()
	l.Delete("a")
	l.Print()


}