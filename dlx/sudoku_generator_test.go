package dlx

import (
	"testing"
)

func Test_GeneratePuzzle(t *testing.T) {
	squareLength := 3
	minSubGivens := 0
	maxSubGivens := squareLength * squareLength
	var puzzle string = GeneratePuzzle(squareLength, minSubGivens, maxSubGivens)
	printSudokuByRaw(squareLength, puzzle)
}

func Test_randomTerminalPattern(t *testing.T) {
	squareLength := 3
	p := newPuzzle(squareLength)
	printSudokuByDigits(squareLength, p.randomTerminalPattern())
}