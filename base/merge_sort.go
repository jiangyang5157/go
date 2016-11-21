package main

import (
	"fmt"
	"github.com/jiangyang5157/go/utils"
)

func merge(left, right []int) []int {
	leftLen := len(left)
	rightLen := len(right)
	ret := make([]int, 0, leftLen + rightLen)

	for ;leftLen > 0 || rightLen > 0; leftLen, rightLen = len(left), len(right) {
		if leftLen == 0 {
			return append(ret, right...)
		}
		if rightLen == 0 {
			return append(ret, left...)
		}

		if right[0] < left[0] {
			ret = append(ret, right[0])
			right = right[1:]
		} else {
			ret = append(ret, left[0])
			left = left[1:]
		}
	}

	return ret
}

// O(n log n)
func mergeSort(arr []int) []int {
	len := len(arr)
	if len <= 1 {
		return arr
	}

	middleIndex := len / 2

	leftPart := mergeSort(arr[:middleIndex])
	rightPart := mergeSort(arr[middleIndex:])

	return merge(leftPart, rightPart)
}

func main() {
	arr := utils.RandomArray(10)
	fmt.Println("Initial array is:", arr)
	fmt.Println("Sorted array is: ", mergeSort(arr))
}
