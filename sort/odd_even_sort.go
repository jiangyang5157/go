package sort

// O(n^2)
func OddEvenSort(arr []int) []int {
	arrLen := len(arr)
	for isSorted := false; isSorted == false; {
		isSorted = true
		for i := 0; i < arrLen-1; i += 2 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isSorted = false
			}
		}
		for i := 1; i < arrLen-1; i += 2 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isSorted = false
			}
		}
	}
	return arr
}
