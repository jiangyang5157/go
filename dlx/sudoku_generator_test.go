package dlx

import (
	"testing"
)

func Test_GeneratePuzzle(t *testing.T) {
	squareLength := 3
	minTotalGivens := 0
	minSubGivens := 0
	maxSubGivens := squareLength * squareLength
	var puzzle string = GeneratePuzzle(squareLength, minTotalGivens, minSubGivens, maxSubGivens)
	printSudokuByRaw(squareLength, puzzle)
}

func Test_randomTerminalPattern(t *testing.T) {
	squareLength := 3
	p := newPuzzle(squareLength)
	printSudokuByDigits(squareLength, p.randomTerminalPattern())
}