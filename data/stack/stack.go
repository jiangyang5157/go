package stack

type Stack struct {
	data   []interface{}
	length int
}

func New() *Stack {
	return &Stack{data: make([]interface{}, 0), length: 0}
}

func (stack *Stack) Length() int {
	return stack.length
}

func (stack *Stack) IsEmpty() bool {
	return stack.length == 0
}

func (stack *Stack) Peek() interface{} {
	return stack.data[stack.length - 1]
}

func (stack *Stack) Pop() interface{} {
	tmp := stack.Peek()
	stack.data = stack.data[0:stack.length]
	stack.length--
	return tmp
}

func (stack *Stack) Push(data interface{}) {
	stack.data = append(stack.data, data)
	stack.length++
}