package queue

type Queue struct {
	data []interface{}
	len  int
}

func New() *Queue {
	queue := &Queue{}
	queue.data = make([]interface{}, 0)
	queue.len = 0
	return queue
}

func (queue *Queue) Length() int {
	return queue.len
}

func (queue *Queue) IsEmpty() bool {
	return queue.len == 0
}

func (queue *Queue) Peek() interface{} {
	return queue.data[0]
}

func (queue *Queue) Remove() interface{} {
	tmp := queue.Peek()
	queue.data = queue.data[1:]
	queue.len -= 1
	return tmp
}

func (queue *Queue) Insert(data interface{}) {
	queue.data = append(queue.data, data)
	queue.len += 1
}