package _8_stack

import "testing"

func TestStack_Push(t *testing.T) {
	stack := New(10)
	for i := 0; i < 12 ; i++  {
		stack.Push("abc")
	}
}

func TestStack_Pop(t *testing.T) {
	stack := New(10)
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
