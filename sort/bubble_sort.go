package sort

// O(n) already sorted; O(n^2); O(n^2)
func BubbleSort(arr []int) []int {
	arrLen := len(arr)
	for i := 0; i < arrLen; i++ {
		for j := 0; j < arrLen-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
