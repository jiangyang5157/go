package set

type Iterator struct {
	ch   <-chan interface{}
	stop chan struct{}
}

func newIterator() (*Iterator, chan<- interface{}, <-chan struct{}) {
	itemChan := make(chan interface{})
	stopChan := make(chan struct{})
	return &Iterator{
		ch:   itemChan,
		stop: stopChan,
	}, itemChan, stopChan
}

func (iter *Iterator) Stop() {
	defer func() {
		recover()
	}()
	close(iter.stop)
	for _ = range iter.ch {

	}
}
