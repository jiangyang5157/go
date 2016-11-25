package number_test

import (
	"testing"
	"github.com/jiangyang5157/go/number"
	"fmt"
	"github.com/jiangyang5157/go/sort"
)

func Test_Sort(t *testing.T) {
	arr := number.RandomArray(10)
	fmt.Printf("%v as the unsorted array\n", arr)

	fmt.Printf("%v from BubbleSort\n", sort.BubbleSort(arr))
	fmt.Printf("%v from CocktailSort\n", sort.CocktailSort(arr))
	fmt.Printf("%v from CombSort\n", sort.CombSort(arr))
	fmt.Printf("%v from CountingSort\n", sort.CountingSort(arr))
	fmt.Printf("%v from GnomeSort\n", sort.GnomeSort(arr))
	fmt.Printf("%v from InsertSort\n", sort.InsertSort(arr))
	fmt.Printf("%v from MergeSort\n", sort.MergeSort(arr))
	fmt.Printf("%v from OddEvenSort\n", sort.OddEvenSort(arr))
	fmt.Printf("%v from QuickSort\n", sort.QuickSort(arr))
	fmt.Printf("%v from SelectionSort\n", sort.SelectionSort(arr))
	fmt.Printf("%v from ShellSort\n", sort.ShellSort(arr))
}