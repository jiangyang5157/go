package main

import (
	"github.com/jiangyang5157/go/utils"
	"fmt"
)

func getK(arr []int) int {
	arrLen := len(arr);
	if arrLen == 0 {
		return 1
	}

	k := arr[0]
	for _, item := range arr {
		if item > k {
			k = item
		}
	}
	return k + 1
}

// O(n + k) where k is the range of numbers and n is the input size
func countingSort(arr []int) []int {
	k := getK(arr)
	counts := make([]int, k)

	arrLen := len(arr)
	for i := 0; i < arrLen; i++ {
		counts[arr[i]] += 1
	}

	for i, j := 0, 0; i < k; i++ {
		for {
			if counts[i] > 0 {
				counts[i] -= 1
				arr[j] = i
				j += 1
			} else {
				break
			}
		}
	}
	return arr
}

func main() {
	arr := utils.RandomArray(10)
	fmt.Println("Initial array is:", arr)
	fmt.Println("Sorted array is: ", countingSort(arr))
}
