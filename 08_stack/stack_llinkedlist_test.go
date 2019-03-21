package _8_stack

import "testing"

func TestLinkedStack_Push(t *testing.T) {
	stack := NewLinkedStack()
	for i := 0; i < 12 ; i++  {
		stack.Push("abc")
	}
}

func TestLinkedStack_Pop(t *testing.T) {
	stack := NewLinkedStack()
	stack.Push("abc")
	stack.Push("123")
	v1 := stack.Pop()
	v2 := stack.Pop()
	if v1 != "123" {
		t.Error("出错了")
	}
	if v2 != "abc" {
		t.Error("出错了")
	}
}
