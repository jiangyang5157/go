package queue

type Queue struct {
	value []interface{}
	len  int
}

func New() *Queue {
	queue := &Queue{}
	queue.value = make([]interface{}, 0)
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
	return queue.value[0]
}

func (queue *Queue) Remove() interface{} {
	tmp := queue.Peek()
	queue.value = queue.value[1:]
	queue.len -= 1
	return tmp
}

func (queue *Queue) Insert(value interface{}) {
	queue.value = append(queue.value, value)
	queue.len += 1
}