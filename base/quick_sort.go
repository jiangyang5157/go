package main

import (
	"fmt"
	"github.com/jiangyang5157/go/utils"
	"math/rand"
)

// O(n log n) -> O(n^2)
func quickSort(arr []int) []int {
	arrLen := len(arr)
	if len <= 1 {
		return arr
	}

	// Avoid O(n^2) worst case
	median := arr[rand.Intn(len)]

	lowerPart := make([]int, 0, arrLen)
	middlePart := make([]int, 0, arrLen)
	higherPart := make([]int, 0, arrLen)

	// skip index, require value only
	for _, item := range arr {
		switch {
		case item < median:
			lowerPart = append(lowerPart, item)
		case item == median:
			middlePart = append(middlePart, item)
		case item > median:
			higherPart = append(higherPart, item)
		}
	}

	lowerPart, higherPart  = quickSort(lowerPart), quickSort(higherPart)

	lowerPart = append(lowerPart, middlePart...)
	lowerPart = append(lowerPart, higherPart...)
	return lowerPart
}

func main() {
	arr := utils.RandomArray(10)
	fmt.Println("Initial array is:", arr)
	fmt.Println("Sorted array is: ", quickSort(arr))
}