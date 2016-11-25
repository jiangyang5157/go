package number

import "math/rand"

func RandomArray(n int) []int {
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		ret[i] = rand.Intn(n)
	}
	return ret
}
