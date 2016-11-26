package stack

type Stack struct {
	value []interface{}
	len  int
}

func New() *Stack {
	stack := &Stack{}
	stack.value = make([]interface{}, 0)
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
	return stack.value[stack.len - 1]
}

func (stack *Stack) Pop() interface{} {
	tmp := stack.Peek()
	stack.value = stack.value[0:stack.len]
	stack.len -= 1
	return tmp
}

func (stack *Stack) Push(value interface{}) {
	stack.value = append(stack.value, value)
	stack.len += 1
}