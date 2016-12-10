package search

// O(log n) always
func BinarySearch(arr []int, find int) int {
	low, high := arr[0], arr[len(arr)-1]
	if find < low || high < find {
		return -1
	}

	for low <= high {
		mid := (low + high) / 2
		switch {
		case arr[mid] < find:
			low = mid + 1
		case arr[mid] > find:
			high = mid - 1
		case arr[mid] == find:
			return mid
		}
	}
	return -1
}
