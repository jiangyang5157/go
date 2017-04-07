package tries

import (
	"fmt"
	"testing"
)

func Test_Add(t *testing.T) {
	root := NewTries()
	fmt.Println(root)
	root.Add("abAB")
	root.Add("abABCD")
	root.Add("bc")
	root.Add("a")
	root.Add("bc")
	fmt.Println(len(root.children))
}

func Test_Search(t *testing.T) {
	root := NewTries()
	root.Add("abAB")
	root.Add("abABCD")
	root.Add("bc")
	root.Add("a")
	root.Add("bc")
	fmt.Println("Test_Search")
	fmt.Println(root.Search("asdasdasdasd"))
	fmt.Println(root.Search("abABC"))
	fmt.Println(root.Search("abABCD"))
	fmt.Println(root.Search(""))
	fmt.Println(root.SearchBeginWith("asdasdasdasd"))
	fmt.Println(root.SearchBeginWith("a"))
	fmt.Println(root.SearchBeginWith("ab"))
	fmt.Println(root.SearchBeginWith(""))
}

func Test_Remove(t *testing.T) {
	root := NewTries()
	root.Add("abAB")
	root.Add("abABCD")
	root.Add("bc")
	root.Add("a")
	root.Add("bc")
	fmt.Println(root.Remove("asdasdasdasd"))
	fmt.Println(root.Remove("abABC"))
	fmt.Println(root.Remove("a"))
	fmt.Println(root.Remove("ab"))
	fmt.Println(root.Remove("abABCD"))
	fmt.Println(root.Remove("abAB"))
	root.Add("abABCD")
	root.Add("abABCD")
}
