package main

import (
	"fmt"
	"github.com/jiangyang5157/go/number"
)

// O(n^2) -> O(n) if the list is initially almost sorted
func gnomeSort(arr []int) []int {
	arrLen := len(arr)
	for i := 1; i < arrLen; {
		if arr[i - 1] > arr[i] {
			arr[i - 1], arr[i] = arr[i], arr[i - 1]
			if i > 1 {
				i--
			}
		} else {
			i++
		}
	}
	return arr
}

func main() {
	arr := number.RandomArray(10)
	fmt.Println("Unsorted array is: ", arr)
	fmt.Println("Sorted array is: ", gnomeSort(arr))
}
