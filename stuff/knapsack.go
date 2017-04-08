package stuff

//  Given a set of items, each with a weight and a value, determine the number of each item to include in a collection so that the total weight is less than or equal to a given limit and the total value is as large as possible.

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

func knapsack() {
	const c = 10                       // capacity
	const n = 5                        // number of items
	v := [n + 1]int{0, 2, 4, 8, 10, 6} // value of items
	w := [n + 1]int{0, 1, 2, 4, 5, 3}  // weight of items
	m := [n + 1][c + 1]int{}

	// Memoization: storing precomputed values
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
