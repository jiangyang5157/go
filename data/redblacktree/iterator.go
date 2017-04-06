package redblacktree

type Iterator func(KeyType) bool

// Ascend will call iterator once for each element greater or equal than pivot
// in ascending order. It will stop whenever the iterator returns false.
func (t *RbTree) Ascend(pivot KeyType, iterator Iterator) {
	t.ascend(t.root, pivot, iterator)
}

func (t *RbTree) ascend(x *Node, pivot KeyType, iterator Iterator) bool {
	// if x == t.NIL {
	// 	return true
	// }

	if !lessThan(x.Key, pivot) {
		if !t.ascend(x.left, pivot, iterator) {
			return false
		}
		if !iterator(x.Key) {
			return false
		}
	}

	return t.ascend(x.right, pivot, iterator)
}

// Descend will call iterator once for each element less or equal than pivot
// in descending order. It will stop whenever the iterator returns false.
func (t *RbTree) Descend(pivot KeyType, iterator Iterator) {
	t.descend(t.root, pivot, iterator)
}

func (t *RbTree) descend(x *Node, pivot KeyType, iterator Iterator) bool {
	// if x == t.NIL {
	// 	return true
	// }

	if !lessThan(pivot, x.Key) {
		if !t.descend(x.right, pivot, iterator) {
			return false
		}
		if !iterator(x.Key) {
			return false
		}
	}

	return t.descend(x.left, pivot, iterator)
}

// AscendRange will call iterator once for elements greater or equal than @ge
// and less than @lt, which means the range would be [ge, lt).
// It will stop whenever the iterator returns false.
func (t *RbTree) AscendRange(ge, lt KeyType, iterator Iterator) {
	t.ascendRange(t.root, ge, lt, iterator)
}

func (t *RbTree) ascendRange(x *Node, inf, sup KeyType, iterator Iterator) bool {
	// if x == t.NIL {
	// 	return true
	// }

	if !lessThan(x.Key, sup) {
		return t.ascendRange(x.left, inf, sup, iterator)
	}
	if lessThan(x.Key, inf) {
		return t.ascendRange(x.right, inf, sup, iterator)
	}

	if !t.ascendRange(x.left, inf, sup, iterator) {
		return false
	}
	if !iterator(x.Key) {
		return false
	}
	return t.ascendRange(x.right, inf, sup, iterator)
}
