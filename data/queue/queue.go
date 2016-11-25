package queue

type Queue struct {
	data []interface{}
	len  int
}

func New() *Queue {
	q := &Queue{}
	q.data = make([]interface{}, 0)
	q.len = 0
	return q
}

func (q *Queue) Length() int {
	return q.len
}

func (q *Queue) Add(data interface{}) {
	q.data = append(q.data, data)
	q.len += 1
}

func (q *Queue) Remove() interface{} {
	tmp := q.data[0]
	q.data = q.data[1:]
	q.len -= 1
	return tmp
}

func (q *Queue) Peek() interface{} {
	return q.data[0]
}