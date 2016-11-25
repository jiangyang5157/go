package search

func LinearSearch(arr []int, find int) int {
	for i, item := range arr {
		if item == find {
			return i
		}
	}
	return -1
}
