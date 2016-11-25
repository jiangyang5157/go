package main

import (
	"fmt"
	"github.com/jiangyang5157/go/number"
)

// O(n^2)
func bubbleSort(arr []int) []int {
	arrLen := len(arr)
	for i := 0; i < arrLen; i++ {
		for j := 0; j < arrLen - 1; j++ {
			if arr[j] > arr[j + 1] {
				arr[j], arr[j + 1] = arr[j + 1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	arr := number.RandomArray(10)
	fmt.Println("Unsorted array is: ", arr)
	fmt.Println("Sorted array is: ", bubbleSort(arr))
}
