package stack

type Stack struct {
	data []interface{}
	len  int
}

func New() *Stack {
	stack := &Stack{}
	stack.len = 0
	stack.data = make([]interface{}, 0)
	return stack
}

func (stack *Stack) Length() int {
	return stack.len
}

func (stack *Stack) IsEmpty() bool {
	return stack.len == 0
}

func (stack *Stack) Peek() interface{} {
	return stack.data[stack.len - 1]
}

func (stack *Stack) Pop() interface{} {
	tmp := stack.Peek()
	stack.len -= 1
	stack.data = stack.data[0:stack.len]
	return tmp
}

func (stack *Stack) Push(data interface{}) {
	stack.len += 1
	stack.data = append(stack.data, data)
}