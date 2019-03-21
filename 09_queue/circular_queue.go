package _9_queue

type CircularQueue struct {
	data []interface{}
	tail uint
	head uint
}

func NewCircularQueue( capacity uint) *CircularQueue {
	return &CircularQueue{
		make([]interface{},capacity),
		0,
		0,
		}
}

func (queue *CircularQueue) Enqueue(data interface{}) bool {
	n := uint(cap(queue.data))
	if i := (queue.tail + 1) % n ; i == queue.head {//队列满了
		return false
	}
	queue.data[queue.tail] = data
	queue.tail = (queue.tail + 1) % n
	return true
}

func (queue *CircularQueue) Dequeue() interface{}  {
	n := uint(cap(queue.data))
	if queue.head == queue.tail {//队列空了
		return nil
	}

	dat := queue.data[queue.head]
	queue.head = (queue.head + 1) % n

	return dat
}