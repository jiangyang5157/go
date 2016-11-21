package main

import (
	"fmt"
	"github.com/jiangyang5157/go/utils"
)

// O(n^2)
func cocktailSort(arr []int) []int {
	for arrLen, i := len(arr), 0; i < arrLen / 2; i++ {
		for left, right := 0, arrLen - 1; left <= right; left, right = left + 1, right - 1 {
			if arr[left] > arr[left + 1] {
				arr[left], arr[left + 1] = arr[left + 1], arr[left]
			}
			if arr[right - 1] > arr[right] {
				arr[right - 1], arr[right] = arr[right], arr[right - 1]
			}
		}
	}
	return arr
}

func main() {
	arr := utils.RandomArray(10)
	fmt.Println("Initial array is:", arr)
	fmt.Println("Sorted array is: ", cocktailSort(arr))
}
