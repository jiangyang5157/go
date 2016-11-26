package queue

type Queue struct {
	data   []interface{}
	length int
}

func New() *Queue {
	return &Queue{data: make([]interface{}, 0), length: 0}
}

func (queue *Queue) Length() int {
	return queue.length
}

func (queue *Queue) IsEmpty() bool {
	return queue.length == 0
}

func (queue *Queue) Peek() interface{} {
	return queue.data[0]
}

func (queue *Queue) Pop() interface{} {
	tmp := queue.Peek()
	queue.data = queue.data[1:]
	queue.length--
	return tmp
}

func (queue *Queue) Push(data interface{}) {
	queue.data = append(queue.data, data)
	queue.length++
}