package heap

/*
Heap is good at find Min/Max O(1)
Heap would take O(n) time to search
*/

type CompareFunc func(a interface{}, b interface{}) int

type Heap struct {
	data    []interface{}
	compare CompareFunc
}

func New(compare CompareFunc) *Heap {
	return &Heap{data: make([]interface{}, 0), compare: compare}
}

func (h *Heap) IsEmpty() bool {
	return h.Length() == 0
}

func (h *Heap) Length() int {
	return len(h.data)
}

func (h *Heap) Get(index int) interface{} {
	return h.data[index]
}

func (h *Heap) siftUp() {
	length := h.Length()
	var parent int
	for child := length - 1; child > 0; {
		// parent: (n - 1) / 2
		parent = (child - 1) >> 1
		if h.compare(h.Get(parent), h.Get(child)) < 0 {
			// if parent "less" then the child, swap
			h.data[parent], h.data[child] = h.data[child], h.data[parent]
			child = parent
		} else {
			break
		}
	}
}

func (h *Heap) siftDown() {
	length := h.Length()
	var child int
	for parent := 0; parent < length && (parent<<1)+1 < length; {
		// left: 2n + 1, right: 2n + 2
		child = (parent << 1) + 1
		if child+1 < length && h.compare(h.Get(child), h.Get(child+1)) < 0 {
			// left child "less" then the right child, position to right child
			child++
		}
		if h.compare(h.Get(parent), h.Get(child)) < 0 {
			// if parent "less" then the "greatest" child, swap
			h.data[parent], h.data[child] = h.data[child], h.data[parent]
			parent = child
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
// O(1)
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
// O(1)
func (heap *Heap) Peek() interface{} {
	if heap.Length() == 0 {
		return nil
	}
	return heap.data[0]
}
