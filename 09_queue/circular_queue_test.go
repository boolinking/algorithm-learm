package _9_queue

import "testing"

func TestCircularQueue_Enqueue(t *testing.T) {
	queue := NewCircularQueue(10)
	queue.Enqueue("abc")
	queue.Enqueue("123")
	queue.Enqueue(123)
	v1 := queue.Dequeue()
	v2 := queue.Dequeue()
	v3 := queue.Dequeue()
	if v1 != "abc" {
		t.Error("出错了")
	}
	if v2 != "123" {
		t.Error("出错了")
	}
	if v3 != 123 {
		t.Error("出错了")
	}
	queue.Enqueue("ab")
	v4 := queue.Dequeue()
	if v4 != "ab" {
		t.Error("出错了")
	}
}

func TestCircularQueue_Dequeue(t *testing.T) {
	queue := NewCircularQueue(10)
	for i := 0; i < 10 ; i++  {
		queue.Enqueue("abc")
	}
	for i := 0; i < 10 ; i++  {
		queue.Dequeue()
	}

	queue.Enqueue("123")
	v := queue.Dequeue()
	if v != "123" {
		t.Error("出错了")
	}
}
