package main

import (
	"fmt"
	"github.com/jiangyang5157/go/number"
)

// O(n log n) best case -> O(n) worst-case
func shellSort(arr []int) []int {
	arrLen := len(arr);
	for gap := int(arrLen / 2); gap > 0; gap /= 2 {
		for i := gap; i < arrLen; i++ {
			for j := i; j >= gap && arr[j - gap] > arr[j]; j -= gap {
				arr[j], arr[j - gap] = arr[j - gap], arr[j]
			}
		}
	}
	return arr
}

func main() {
	arr := number.RandomArray(10)
	fmt.Println("Initial array is:", arr)
	fmt.Println("Sorted array is: ", shellSort(arr))
}
