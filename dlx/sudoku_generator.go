package dlx

import (
	"time"
	"math/rand"
)

// Generate Sudoku puzzle that has unique solution
func GenerateSudoku(squareLength int, minTotalGivens int, minSubGivens int, maxSubGivens int) string {
	// todo
	return ""
}

func randomArray(length int, min int, max int) *[]int {
	var ret []int = make([]int, length)
	valueRange := max - min
	switch {
	case valueRange < 0:
		return nil
	case valueRange == 0:
		for i := 0; i < length; i++ {
			ret[i] = min
		}
		return &ret
	default:
		rand.Seed(time.Now().Unix())
		for i := 0; i < length; i++ {
			ret[i] = rand.Intn(valueRange + 1) + min
		}
		return &ret
	}
}