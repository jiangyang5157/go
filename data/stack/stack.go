package stack

type Stack struct {
	value []interface{}
	size  int
}

func New() *Stack {
	return &Stack{value: make([]interface{}, 0), size: 0}
}

func (stack *Stack) Size() int {
	return stack.size
}

func (stack *Stack) IsEmpty() bool {
	return stack.size == 0
}

func (stack *Stack) Peek() interface{} {
	return stack.value[stack.size - 1]
}

func (stack *Stack) Pop() interface{} {
	tmp := stack.Peek()
	stack.value = stack.value[0:stack.size]
	stack.size--
	return tmp
}

func (stack *Stack) Push(value interface{}) {
	stack.value = append(stack.value, value)
	stack.size++
}