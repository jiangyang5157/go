package string

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// Longest common subsequence recurrence
func Lcs(x string, y string) string {
	m, n := len(x), len(y)

	opt := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		opt[i] = make([]int, n+1)
	}

	// Compute length of LCS for all sub-problems
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if x[i] == y[j] {
				opt[i][j] = opt[i+1][j+1] + 1
			} else {
				opt[i][j] = max(opt[i+1][j], opt[i][j+1])
			}
		}
	}

	// Recover LCS itself
	lcs := ""
	for i, j := 0, 0; i < m && j < n; {
		if x[i] == y[j] {
			lcs += string(x[i])
			i++
			j++
		} else {
			if opt[i+1][j] >= opt[i][j+1] {
				i++
			} else {
				j++
			}
		}
	}

	return lcs
}
