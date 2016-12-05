package heap

type Comparable func(a interface{}, b interface{}) int

type Heap struct {
	data    []interface{}
	compare Comparable
}

func New(compare Comparable) *Heap {
	return &Heap{data: make([]interface{}, 0), compare: compare}
}

func (heap *Heap) IsEmpty() bool {
	return heap.Length() == 0
}

func (heap *Heap) Length() int {
	return len(heap.data)
}

func (heap *Heap) Get(index int) interface{} {
	return heap.data[index]
}

func (heap *Heap) siftUp() {
	length := heap.Length()
	for child, parent := length-1, length-1; child > 0; child = parent {
		// parent: (n - 1) / 2
		parent = (child - 1) >> 1
		if heap.compare(heap.Get(parent), heap.Get(child)) < 0 {
			// if parent "less" then the child, swap
			heap.data[parent], heap.data[child] = heap.data[child], heap.data[parent]
		} else {
			break
		}
	}
}

func (heap *Heap) siftDown() {
	length := heap.Length()
	for parent, child := 0, 1; parent < length && (parent<<1)+1 < length; parent = child {
		// left: 2n + 1, right: 2n + 2
		child = (parent << 1) + 1
		if child+1 < length && heap.compare(heap.Get(child), heap.Get(child+1)) < 0 {
			// left child "less" then the right child, position to right child
			child++
		}
		if heap.compare(heap.Get(parent), heap.Get(child)) < 0 {
			// if parent "less" then the "greatest" child, swap
			heap.data[parent], heap.data[child] = heap.data[child], heap.data[parent]
		} else {
			break
		}
	}
}

// Insert element into the next position
func (heap *Heap) Insert(data interface{}) {
	heap.data = append(heap.data, data)

	heap.siftUp()
}

// Extract the root element
func (heap *Heap) Extract() interface{} {
	length := heap.Length()
	if length == 0 {
		return nil
	}
	data := heap.data[0]
	// replace the root with the last element
	heap.data[0] = heap.data[length-1]
	heap.data = heap.data[:length-1]

	heap.siftDown()
	return data
}

// Peek at the root
func (heap *Heap) Peek() interface{} {
	if heap.Length() == 0 {
		return nil
	}
	return heap.data[0]
}
