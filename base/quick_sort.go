package main

import (
	"fmt"
	"github.com/jiangyang5157/go/utils"
	"math/rand"
)

func quickSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}

	// Avoid O(n^2) worst case
	median := arr[rand.Intn(length)]

	lower := make([]int, 0, length)
	middle := make([]int, 0, length)
	higher := make([]int, 0, length)

	// skip index, require value only
	for _, item := range arr {
		switch {
		case item < median:
			lower = append(lower, item)
		case item == median:
			middle = append(middle, item)
		case item > median:
			higher = append(higher, item)
		}
	}

	lower = quickSort(lower)
	higher = quickSort(higher)

	lower = append(lower, middle...)
	lower = append(lower, higher...)

	return lower
}

func main() {
	arr := utils.RandomArray(10)
	fmt.Println("Initial array is:", arr)
	fmt.Println("Sorted array is: ", quickSort(arr))
}