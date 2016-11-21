package main

import (
	"fmt"
	"github.com/jiangyang5157/go/utils"
)

// O(n^2)
func selectionSort(arr []int) []int {
	for arrLen, min, i := len(arr), 0, 0; i < arrLen - 1; i++ {
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
	arr := utils.RandomArray(10)
	fmt.Println("Initial array is:", arr)
	fmt.Println("Sorted array is: ", selectionSort(arr))
}
