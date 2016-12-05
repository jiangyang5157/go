package sort

// O(n^2) -> O(n) if the list is initially almost sorted
func GnomeSort(arr []int) []int {
	arrLen := len(arr)
	for i := 1; i < arrLen; {
		if arr[i-1] > arr[i] {
			arr[i-1], arr[i] = arr[i], arr[i-1]
			if i > 1 {
				i--
			}
		} else {
			i++
		}
	}
	return arr
}
