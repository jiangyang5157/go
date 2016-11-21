package main

import (
	"fmt"
	"github.com/jiangyang5157/go/utils"
)

// O(n^2)
func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr) - 1; j++ {
			if arr[j] > arr[j + 1] {
				tmp := arr[j]
				arr[j] = arr[j + 1]
				arr[j + 1] = tmp
			}
		}
	}
	return arr
}

func main() {
	arr := utils.RandomArray(10)
	fmt.Println("Initial array is:", arr)
	fmt.Println("Sorted array is: ", bubbleSort(arr))
}
