package queue

type Queue struct {
	data []interface{}
	len  int
}

func New() *Queue {
	queue := &Queue{}
	queue.len = 0
	queue.data = make([]interface{}, 0)
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
	queue.len -= 1
	queue.data = queue.data[1:]
	return tmp
}

func (queue *Queue) Add(data interface{}) {
	queue.len += 1
	queue.data = append(queue.data, data)
}