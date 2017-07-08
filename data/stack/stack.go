package stack

type Stack struct {
	data   []interface{}
	length int
}

func NewStack() *Stack {
	return &Stack{data: make([]interface{}, 0), length: 0}
}

func (s *Stack) Length() int {
	return s.length
}

func (s *Stack) IsEmpty() bool {
	return s.Length() == 0
}

func (s *Stack) Peek() interface{} {
	return s.data[s.length-1]
}

func (s *Stack) Pop() interface{} {
	tmp := s.Peek()
	s.data = s.data[:s.length-1]
	s.length--
	return tmp
}

func (s *Stack) Push(data interface{}) {
	s.data = append(s.data, data)
	s.length++
}
