package queue

type Queue struct {
	data   []interface{}
	length int
}

func NewQueue() *Queue {
	return &Queue{data: make([]interface{}, 0), length: 0}
}

func (q *Queue) Length() int {
	return q.length
}

func (q *Queue) IsEmpty() bool {
	return q.Length() == 0
}

func (q *Queue) Peek() interface{} {
	return q.data[0]
}

func (q *Queue) Pop() interface{} {
	if q.IsEmpty() {
		return nil
	}
	ret := q.Peek()
	q.data = q.data[1:]
	q.length--
	return ret
}

func (q *Queue) Push(data interface{}) {
	q.data = append(q.data, data)
	q.length++
}
