package number_test

import (
	"testing"
	"github.com/jiangyang5157/go/number"
	"fmt"
	"github.com/jiangyang5157/go/sort"
	"github.com/jiangyang5157/go/search"
)

func Test_Sort(t *testing.T) {
	arr := sort.MergeSort(number.RandomArray(10));
	canBefind, canNotBefind := arr[3], 11;
	fmt.Printf("Find %v and %v in the sorted arrya %v\n", canBefind, canNotBefind, arr)

	fmt.Printf("Found index %v by LinearSearch\n", search.LinearSearch(arr, canBefind))
	fmt.Printf("Found index %v by LinearSearch\n", search.LinearSearch(arr, canNotBefind))
	fmt.Printf("Found index %v by BinarySearch\n", search.BinarySearch(arr, canBefind))
	fmt.Printf("Found index %v by BinarySearch\n", search.BinarySearch(arr, canNotBefind))
}