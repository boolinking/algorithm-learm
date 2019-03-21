package _9_queue

type ArrayQueue struct {
	data []interface{}
	tail uint
	head uint
}

func NewArrayQueue(capacity uint) *ArrayQueue {
	return &ArrayQueue{
		make([]interface{},capacity),
		0,
		0,
	}
}

func (queue *ArrayQueue) Enqueue(data interface{}) bool {
	if queue.tail == uint(cap(queue.data)) {

		if queue.head == 0 { //队列满了
			return false
		}

		for i := queue.head; i < queue.tail ; i++ {//数据迁移
			queue.data[i-queue.head] = queue.data[i]
		}
		queue.tail -= queue.head
		queue.head = 0
	}
	queue.data[queue.tail] = data
	queue.tail++
	return true
}

func (queue *ArrayQueue) Dequeue() interface{}  {
	if queue.head == queue.tail {
		return nil
	}
	v := queue.data[queue.head]
	queue.head++
	return v
}
