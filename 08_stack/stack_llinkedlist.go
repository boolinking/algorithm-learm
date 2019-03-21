package _8_stack

type Node struct {
	next *Node
	data interface{}
}


type LinkedStack struct {
	head *Node
	length uint
}

func NewLinkedStack() *LinkedStack {
	return &LinkedStack{
		nil,0,
	}
}



func (stack *LinkedStack) Push(data interface{}) {
	n := &Node{nil,data}
	n.next = stack.head
	stack.head = n
	stack.length++
}

func (stack *LinkedStack) Pop() interface{}  {
	if stack.head == nil {
		return nil
	}
	data := stack.head.data
	stack.head = stack.head.next
	stack.length--
	return  data
}

func (stack *LinkedStack) Clear()  {
	stack.head = nil
	stack.length = 0
}