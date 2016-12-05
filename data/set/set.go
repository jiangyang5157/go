package set

type Set interface {
	String() string
	Equals(other Set) bool
	Size() int
	Clear()
	Clone() Set
	//Returns whether the item was added.
	Add(i interface{}) bool
	// Remove a single element from the set.
	Remove(i interface{})
	// Returns a new set with all elements in both sets.
	Union(other Set) Set
	Contains(i ...interface{}) bool
	// Returns a new set containing only the elements that exist only in both sets.
	Intersect(other Set) Set
	// Determines if every element in this set is in the other set.
	IsSubset(other Set) bool
	// Determines if every element in the other set is in this set.
	IsSuperset(other Set) bool
	// The returned set will contain all elements of this set that are not also elements of other.
	Difference(other Set) Set
	// Returns a new set with all elements which are in either this set or the other set but not in both.
	SymmetricDifference(other Set) Set
	// Returns an Iterator object that you can use to range over the set.
	Iterator() *Iterator
}

func NewThreadUnsafeSet(i ...interface{}) Set {
	set := newThreadUnsafeSet()
	for _, e := range i {
		set.Add(e)
	}
	return &set
}

func NewThreadSafeSet(i ...interface{}) Set {
	set := newThreadSafeSet()
	for _, e := range i {
		set.Add(e)
	}
	return &set
}

func NewThreadUnsafeSetFromSlice(i []interface{}) Set {
	return NewThreadUnsafeSet(i...)
}

func NewThreadSafeSetFromSlice(i []interface{}) Set {
	return NewThreadSafeSet(i...)
}
