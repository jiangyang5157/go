package main

import (
	"fmt"
	"github.com/jiangyang5157/go/number"
)

// O(n^2) average
func selectionSort(arr []int) []int {
	arrLen := len(arr)
	for min, i := 0, 0; i < arrLen - 1; i++ {
		min = i
		for j := i + 1; j < arrLen; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if (min != i) {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
	return arr
}

func main() {
	arr := number.RandomArray(10)
	fmt.Println("Initial array is:", arr)
	fmt.Println("Sorted array is: ", selectionSort(arr))
}
