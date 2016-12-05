package sort

// O(n^2) average
func InsertSort(arr []int) []int {
	arrLen := len(arr)
	for i := 1; i < arrLen; i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
	return arr
}
