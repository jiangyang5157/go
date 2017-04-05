package set

import (
	"fmt"
	"strings"
)

type threadUnsafeSet map[interface{}]struct{}

func newThreadUnsafeSet() threadUnsafeSet {
	return make(threadUnsafeSet)
}

func (set *threadUnsafeSet) String() string {
	pairs := make([]string, 0, set.Size())
	for e := range *set {
		pairs = append(pairs, fmt.Sprintf("%v", e))
	}
	return fmt.Sprintf("Set{%s}", strings.Join(pairs, ", "))
}

func (set *threadUnsafeSet) Equals(other Set) bool {
	if set.Size() != other.Size() {
		return false
	}
	return set.IsSubset(other)
}

func (set *threadUnsafeSet) Size() int {
	return len(*set)
}

func (set *threadUnsafeSet) Clear() {
	*set = newThreadUnsafeSet()
}

func (set *threadUnsafeSet) Clone() Set {
	clone := newThreadUnsafeSet()
	for e := range *set {
		clone.Add(e)
	}
	return &clone
}

func (set *threadUnsafeSet) Add(i interface{}) bool {
	_, found := (*set)[i]
	(*set)[i] = struct{}{}
	return !found
}

func (set *threadUnsafeSet) Remove(i interface{}) {
	delete(*set, i)
}

func (set *threadUnsafeSet) Union(other Set) Set {
	unioned := newThreadUnsafeSet()
	for e := range *set {
		unioned.Add(e)
	}
	o := other.(*threadUnsafeSet)
	for e := range *o {
		unioned.Add(e)
	}
	return &unioned
}

func (set *threadUnsafeSet) Contains(i ...interface{}) bool {
	for _, val := range i {
		if _, found := (*set)[val]; !found {
			return false
		}
	}
	return true
}

func (set *threadUnsafeSet) Intersect(other Set) Set {
	intersection := newThreadUnsafeSet()
	if set.Size() < other.Size() {
		for e := range *set {
			if other.Contains(e) {
				intersection.Add(e)
			}
		}
	} else {
		o := other.(*threadUnsafeSet)
		for e := range *o {
			if set.Contains(e) {
				intersection.Add(e)
			}
		}
	}
	return &intersection
}

func (set *threadUnsafeSet) IsSubset(other Set) bool {
	for e := range *set {
		if !other.Contains(e) {
			return false
		}
	}
	return true
}

func (set *threadUnsafeSet) IsSuperset(other Set) bool {
	return other.IsSubset(set)
}

func (set *threadUnsafeSet) Difference(other Set) Set {
	difference := newThreadUnsafeSet()
	for e := range *set {
		if !other.Contains(e) {
			difference.Add(e)
		}
	}
	return &difference
}

func (set *threadUnsafeSet) SymmetricDifference(other Set) Set {
	a := set.Difference(other)
	b := other.Difference(set)
	return a.Union(b)
}

func (set *threadUnsafeSet) Iterator() *Iterator {
	iterator, itemChan, stopCh := newIterator()
	go func() {
	A:
		for e := range *set {
			select {
			case <-stopCh:
				break A
			case itemChan <- e:
			}
		}
		close(itemChan)
	}()
	return iterator
}
