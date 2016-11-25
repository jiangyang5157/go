package main

import (
	"fmt"
	"github.com/jiangyang5157/go/number"
)

// O(n^2)
func oddEvenSort(arr []int) []int {
	arrLen := len(arr)
	for isSorted := false; isSorted == false; {
		isSorted = true
		for i := 0; i < arrLen - 1; i += 2 {
			if arr[i] > arr[i + 1] {
				arr[i], arr[i + 1] = arr[i + 1], arr[i];
				isSorted = false
			}
		}
		for i := 1; i < arrLen - 1; i += 2 {
			if arr[i] > arr[i + 1] {
				arr[i], arr[i + 1] = arr[i + 1], arr[i];
				isSorted = false
			}
		}
	}
	return arr
}

func main() {
	arr := number.RandomArray(10)
	fmt.Println("Initial array is:", arr)
	fmt.Println("Sorted array is: ", oddEvenSort(arr))
}
