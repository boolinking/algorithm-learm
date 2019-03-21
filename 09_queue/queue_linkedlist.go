package _9_queue

type LinkedListQueue struct {
	head *Node
	tail *Node

}

type Node struct {
	next *Node
	data interface{}
}

func New() *LinkedListQueue {
	return &LinkedListQueue{nil,nil}
}


func (queue *LinkedListQueue) Enqueue(data interface{})  {
	newNode := &Node{nil,data}
	if queue.tail == nil {
		queue.head = newNode
		queue.tail = newNode

	}
	queue.tail.next = newNode
	queue.tail = newNode

}


func (queue *LinkedListQueue) Dequeue() interface{}  {
	if queue.head == nil {
		return nil
	}
	dat := queue.head.data
	queue.head = queue.head.next
	if queue.head == nil {
		queue.tail = nil
	}
	return dat
}

