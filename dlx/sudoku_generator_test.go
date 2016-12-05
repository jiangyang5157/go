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