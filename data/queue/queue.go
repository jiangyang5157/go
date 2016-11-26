package queue

type Queue struct {
	value []interface{}
	size  int
}

func New() *Queue {
	return &Queue{value: make([]interface{}, 0), size: 0}
}

func (queue *Queue) Size() int {
	return queue.size
}

func (queue *Queue) IsEmpty() bool {
	return queue.size == 0
}

func (queue *Queue) Peek() interface{} {
	return queue.value[0]
}

func (queue *Queue) Pop() interface{} {
	tmp := queue.Peek()
	queue.value = queue.value[1:]
	queue.size--
	return tmp
}

func (queue *Queue) Push(value interface{}) {
	queue.value = append(queue.value, value)
	queue.size++
}