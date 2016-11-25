package stack

type Stack struct {
	data []interface{}
	len  int
}

func New() *Stack {
	stack := &Stack{}
	stack.data = make([]interface{}, 0)
	stack.len = 0
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
	stack.data = stack.data[0:stack.len]
	stack.len -= 1
	return tmp
}

func (stack *Stack) Push(data interface{}) {
	stack.data = append(stack.data, data)
	stack.len += 1
}