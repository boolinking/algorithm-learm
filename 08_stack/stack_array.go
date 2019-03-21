package _8_stack

type Stack struct {
	data [] interface{}
	count uint
}

func New(n uint) *Stack {
	return &Stack{
		make([]interface{},n),
		0,
	}
}

func (stack *Stack) Push(data interface{}) bool {
	if stack.count == uint(cap(stack.data)) {//栈满了
		return false
	}
	stack.data[stack.count] = data
	stack.count++
	return true
}

func (stack *Stack) Pop() interface{} {
	if stack.count == 0 {
		return nil
	}
	data := stack.data[stack.count - 1]
	stack.count--
	return data
}