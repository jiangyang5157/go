package list

import (
	"testing"
	"fmt"
)

func Test_List(t *testing.T) {
	list := NewList()

	if list.Size() != 0 {
		t.Error("Size doesn't work as expected")
	}
	if list.IsEmpty() != true {
		t.Error("IsEmpty doesn't work as expected")
	}
	if e, _ := list.Find(222); e != nil {
		t.Error("Find doesn't work as expected")
	}
	if err := list.Remove(222); err == nil {
		t.Error("Remove doesn't work as expected")
	}
	list.Append(222)
	if list.IsEmpty() != false {
		t.Error("IsEmpty doesn't work as expected")
	}
	list.Each(func(e *Element) {
		fmt.Printf("%v, ", e.value)
	})
	fmt.Println("")
	if e, _ := list.Find(222); e == nil {
		t.Error("Find doesn't work as expected")
	}
	if err := list.Remove(222); err != nil {
		t.Error("Remove doesn't work as expected")
	}
	list.Each(func(e *Element) {
		fmt.Printf("%v, ", e.value)
	})
	fmt.Println("")

	list.Prepend(3)
	list.Prepend(2)
	list.Prepend(1)
	list.Append(4)
	list.Append(5)
	list.Append(6)

	if list.Size() != 6 {
		t.Error("Size doesn't work as expected")
	}

	if err := list.Remove(222); err == nil {
		t.Error("Remove doesn't work as expected")
	}
	if list.Size() != 6 {
		t.Error("Size doesn't work as expected")
	}

	list.Each(func(e *Element) {
		fmt.Printf("%v, ", e.value)
	})
	fmt.Println("")
	list.Each(func(e *Element) {
		e.value = 111;
	})
	list.Each(func(e *Element) {
		fmt.Printf("%v, ", e.value)
	})
	fmt.Println("")
	tmp := NewList()
	tmp.Append(222)
	tmp.Append(222)
	tmp.Append(222)
	tmp.Append(222)
	list.Concat(tmp)
	tmp = nil
	list.Each(func(e *Element) {
		fmt.Printf("%v, ", e.value)
	})
	fmt.Println("")
	fmt.Printf("Count 1: %v\n", list.Count(1))
	fmt.Printf("Count 3: %v\n", list.Count(3))
	fmt.Printf("Count 11: %v\n", list.Count(11))
	fmt.Printf("Count 111: %v\n", list.Count(111))
	fmt.Printf("Count 222: %v\n", list.Count(222))
	fmt.Println(list.size)
	list.Clear()
	fmt.Println(list.size)
	list.Each(func(e *Element) {
		fmt.Printf("%v, ", e.value)
	})
	fmt.Println("")
	list.Prepend(3)
	list.Prepend(2)
	list.Prepend(1)
	list.Append(4)
	list.Append(5)
	list.Append(6)
	list.Each(func(e *Element) {
		fmt.Printf("%v, ", e.value)
	})
	fmt.Println("")
	fmt.Println(list.size)
	list.Clear()
	fmt.Println(list.size)
}