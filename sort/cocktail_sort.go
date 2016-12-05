package sort

// O(n^2)
func CocktailSort(arr []int) []int {
	arrLen := len(arr)
	for i := 0; i < arrLen/2; i++ {
		for left, right := 0, arrLen-1; left <= right; left, right = left+1, right-1 {
			if arr[left] > arr[left+1] {
				arr[left], arr[left+1] = arr[left+1], arr[left]
			}
			if arr[right-1] > arr[right] {
				arr[right-1], arr[right] = arr[right], arr[right-1]
			}
		}
	}
	return arr
}
