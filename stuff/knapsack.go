package main

import (
	"fmt"
)

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	const c = 10                       // capacity
	const n = 5                        // number of items
	v := [n + 1]int{0, 2, 4, 8, 10, 6} // value of items
	w := [n + 1]int{0, 1, 2, 4, 5, 3}  // weight of items
	m := [n + 1][c + 1]int{}

	for i := 1; i <= n; i++ {
		for j := 0; j <= c; j++ {
			if w[i] > j {
				m[i][j] = m[i-1][j]
			} else {
				m[i][j] = max(m[i-1][j], m[i-1][j-w[i]]+v[i])
			}
		}
	}

	for i := 0; i <= n; i++ {
		fmt.Println(m[i])
	}
	fmt.Printf("The maximum value of the %v items in %v capacity is %v.\n", n, c, m[n][c])
}
