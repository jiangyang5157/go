package main

import (
	"fmt"
	"github.com/jiangyang5157/go/number"
)

// O(n^2) average
func insertSort(arr []int) []int {
	arrLen := len(arr)
	for i := 1; i < arrLen; i++ {
		for j := i; j > 0 && arr[j] < arr[j - 1]; j-- {
			arr[j], arr[j - 1] = arr[j - 1], arr[j]
		}
	}
	return arr
}

func main() {
	arr := number.RandomArray(10)
	fmt.Println("Initial array is:", arr)
	fmt.Println("Sorted array is: ", insertSort(arr))
}
