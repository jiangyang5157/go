package sort

// O(n^2 / 2^p) average, where p is the number of increments -> O(n^2) worst-case -> O(n log n) best-case
func CombSort(arr []int) []int {
	arrLen := len(arr)
	gap := arrLen
	for {
		if gap > 1 {
			// k = 1.3 has been suggested as an ideal shrink factor by the authors of the original article after empirical testing on over 200,000 random lists
			gap = gap * 100 / 130
		}
		for i := 0; i+gap < arrLen; i++ {
			if arr[i] > arr[i+gap] {
				arr[i], arr[i+gap] = arr[i+gap], arr[i]
			}
		}
		if gap == 1 {
			break
		}
	}
	return arr
}
