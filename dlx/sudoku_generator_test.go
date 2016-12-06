package dlx

import (
	"testing"
	"fmt"
)

func Test_GeneratePuzzle(t *testing.T) {
	squareLength := 3
	minTotalGivens := 0
	minSubGivens := 0
	maxSubGivens := squareLength * squareLength
	fmt.Println(GeneratePuzzle(squareLength, minTotalGivens, minSubGivens, maxSubGivens))
}

func Test_disorderArray(t *testing.T) {
	var digits []int = make([]int, 9)
	for i := range digits {
		digits[i] = i
	}
	fmt.Println(disorderArray(digits))
}