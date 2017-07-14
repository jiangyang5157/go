package sort

import (
	"fmt"
	"os"
	"testing"

	"github.com/jiangyang5157/golang-start/number"
)

var arr []int

func setup() {
	fmt.Println("setup")
	arr = number.RandomArray(10)
	fmt.Printf("%v as the unsorted array\n", arr)
}
func teardown() {
	fmt.Println("teardown")
}

func TestMain(m *testing.M) {
	setup()
	exitCode := m.Run()
	teardown()
	os.Exit(exitCode)
}

func Test_Sort(t *testing.T) {
	fmt.Printf("%v from BubbleSort\n", BubbleSort(arr))
	fmt.Printf("%v from CocktailSort\n", CocktailSort(arr))
	fmt.Printf("%v from CombSort\n", CombSort(arr))
	fmt.Printf("%v from CountingSort\n", CountingSort(arr))
	fmt.Printf("%v from GnomeSort\n", GnomeSort(arr))
	fmt.Printf("%v from InsertSort\n", InsertSort(arr))
	fmt.Printf("%v from MergeSort\n", MergeSort(arr))
	fmt.Printf("%v from OddEvenSort\n", OddEvenSort(arr))
	fmt.Printf("%v from QuickSort\n", QuickSort(arr))
	fmt.Printf("%v from SelectionSort\n", SelectionSort(arr))
	fmt.Printf("%v from ShellSort\n", ShellSort(arr))
}
