package dlx

import (
	"fmt"
	"strings"
	"testing"
)

func solveSudukuTest(squareLength int, raw string, solutionSize int) {
	printSudokuByRaw(squareLength, raw)
	solutions := strings.Split(SolvePuzzleByRaw(squareLength, raw, solutionSize), string(SOLUTION_PREFIX))
	length := len(solutions)
	// [0] contains massage
	fmt.Printf("%v\n", solutions[0])
	fmt.Printf("Looking for %d solutions, found %d solutions.\n", solutionSize, length - 1)
	for i := 1; i < length; i++ {
		printSudokuByRaw(squareLength, solutions[i])
	}
}

func Test_solve(t *testing.T) {
	squareLength := 3
	solutionSize := 5
	raw :=
	//"......123" +
	//"..9......" +
	//".....9..." +
	//"........." +
	//"........." +
	//"........." +
	//"........." +
	//"........." +
	//"........." // 0 solutions puzzle

	//"........." +
	//"..41.26.." +
	//".3..5..2." +
	//".2..1..3." +
	//"..65.41.." +
	//".8..7..4." +
	//".7..2..6." +
	//"..14.35.." +
	//"........." // 1 solutions puzzle

		"..3456789" +
			"456789123" +
			"789123456" +
			"..4365897" +
			"365897214" +
			"897214365" +
			"531642978" +
			"642978531" +
			"978531642" // 2 solutions puzzle

	//"....7.94." +
	//".7..9...5" +
	//"3....5.7." +
	//"..74..1.." +
	//"463.8...." +
	//".....7.8." +
	//"8..7....." +
	//"7......28" +
	//".5..68..." // 188 solutions puzzle

	//"." // 1 solutions puzzle

	//"...." +
	//".4.." +
	//"2..." +
	//"..43" // 0 solutions puzzle

	//"...." +
	//".4.." +
	//"2..." +
	//"...3" // 3 solutions puzzle

	//"................" +
	//"................" +
	//"................" +
	//"................" +
	//"................" +
	//"................" +
	//"................" +
	//"................" +
	//"................" +
	//"................" +
	//"................" +
	//"................" +
	//"................" +
	//"................" +
	//"................" +
	//"................" // ? solutions puzzle

	solveSudukuTest(squareLength, raw, solutionSize)
}