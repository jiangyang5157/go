package dlx

import (
	"testing"
	"strings"
	"fmt"
)

func Test_GeneratePuzzle(t *testing.T) {
	squareLength := 3
	minSubGivens := 2
	maxSubGivens := 6
	var raw string = GeneratePuzzle(squareLength, minSubGivens, maxSubGivens)
	printSudokuByRaw(squareLength, raw)

	p := newPuzzle(squareLength)
	p.init(raw2digits(raw))
	maxSolutionSize := 2
	var solutions []string = strings.Split(p.solvePuzzle(maxSolutionSize), string(SOLUTION_PREFIX))
	fmt.Printf("%v\n", solutions[0])
	fmt.Printf("Looking for %d solutions, found %d solutions.\n", maxSolutionSize, len(solutions) - 1)
	for i := 1; i < len(solutions); i++ {
		printSudokuByRaw(squareLength, solutions[i])
	}
}

func Test_randomTerminalPattern(t *testing.T) {
	squareLength := 3
	p := newPuzzle(squareLength)
	printSudokuByDigits(squareLength, p.randomTerminalPattern())
}