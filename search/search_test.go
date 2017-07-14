package search

import (
	"fmt"
	"testing"

	"github.com/jiangyang5157/golang-start/number"
	"github.com/jiangyang5157/golang-start/sort"
)

func Test_Sort(t *testing.T) {
	arr := sort.MergeSort(number.RandomArray(10))
	canBefind := arr[3]
	canNotBefind := 222
	fmt.Printf("Search %v and %v in the sorted arrya %v\n", canBefind, canNotBefind, arr)

	fmt.Printf("Search %v, index %v by LinearSearch\n", canBefind, LinearSearch(arr, canBefind))
	fmt.Printf("Search %v, index %v by LinearSearch\n", canNotBefind, LinearSearch(arr, canNotBefind))
	fmt.Printf("Search %v, index %v by BinarySearch\n", canBefind, BinarySearch(arr, canBefind))
	fmt.Printf("Search %v, index %v by BinarySearch\n", canNotBefind, BinarySearch(arr, canNotBefind))
}
