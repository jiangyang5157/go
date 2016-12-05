package set

import "sync"

type threadSafeSet struct {
	data threadUnsafeSet
	sync.RWMutex
}

func newThreadSafeSet() threadSafeSet {
	return threadSafeSet{data: newThreadUnsafeSet()}
}

func (set *threadSafeSet) String() string {
	set.RLock()
	defer set.RUnlock()
	return set.data.String()
}

func (set *threadSafeSet) Equals(other Set) bool {
	set.RLock()
	defer set.RUnlock()
	o := other.(*threadSafeSet)
	o.RLock()
	defer o.RUnlock()
	return set.data.Equals(&o.data)
}

func (set *threadSafeSet) Size() int {
	set.RLock()
	defer set.RUnlock()
	return len(set.data)
}

func (set *threadSafeSet) Clear() {
	set.Lock()
	defer set.Unlock()
	set.data = newThreadUnsafeSet()
}

func (set *threadSafeSet) Clone() Set {
	set.RLock()
	defer set.RUnlock()
	unsafeClone := set.data.Clone().(*threadUnsafeSet)
	return &threadSafeSet{data: *unsafeClone}
}

func (set *threadSafeSet) Add(i interface{}) bool {
	set.Lock()
	defer set.Unlock()
	return set.data.Add(i)
}

func (set *threadSafeSet) Remove(i interface{}) {
	set.Lock()
	defer set.Unlock()
	delete(set.data, i)
}

func (set *threadSafeSet) Union(other Set) Set {
	set.RLock()
	defer set.RUnlock()
	o := other.(*threadSafeSet)
	o.RLock()
	defer o.RUnlock()
	unsafeUnion := set.data.Union(&o.data).(*threadUnsafeSet)
	return &threadSafeSet{data: *unsafeUnion}
}

func (set *threadSafeSet) Contains(i ...interface{}) bool {
	set.RLock()
	defer set.RUnlock()
	return set.data.Contains(i...)
}

func (set *threadSafeSet) Intersect(other Set) Set {
	set.RLock()
	defer set.RUnlock()
	o := other.(*threadSafeSet)
	o.RLock()
	defer o.RUnlock()
	unsafeIntersection := set.data.Intersect(&o.data).(*threadUnsafeSet)
	return &threadSafeSet{data: *unsafeIntersection}
}

func (set *threadSafeSet) IsSubset(other Set) bool {
	set.RLock()
	defer set.RUnlock()
	o := other.(*threadSafeSet)
	o.RLock()
	defer o.RUnlock()
	return set.data.IsSubset(&o.data)
}

func (set *threadSafeSet) IsSuperset(other Set) bool {
	return other.IsSubset(set)
}

func (set *threadSafeSet) Difference(other Set) Set {
	set.RLock()
	defer set.RUnlock()
	o := other.(*threadSafeSet)
	o.RLock()
	defer o.RUnlock()
	unsafeDifference := set.data.Difference(&o.data).(*threadUnsafeSet)
	return &threadSafeSet{data: *unsafeDifference}
}

func (set *threadSafeSet) SymmetricDifference(other Set) Set {
	o := other.(*threadSafeSet)
	unsafeDifference := set.data.SymmetricDifference(&o.data).(*threadUnsafeSet)
	return &threadSafeSet{data: *unsafeDifference}
}

func (set *threadSafeSet) Iterator() *Iterator {
	iterator, itemChan, stopCh := newIterator()
	go func() {
		set.RLock()
	A:
		for e := range set.data {
			select {
			case <-stopCh:
				break A
			case itemChan <- e:
			}
		}
		close(itemChan)
		set.RUnlock()
	}()
	return iterator
}
