package dlx

import (
	"fmt"
	"strings"
	"testing"
)

func SolvePuzzleByRaw_Test(squareLength int, raw string, maxSolutionSize int) {
	printSudokuByRaw(squareLength, raw)
	solutions := strings.Split(SolvePuzzleByRaw(squareLength, raw, maxSolutionSize), string(SOLUTION_PREFIX))
	// [0] contains massage
	fmt.Printf("%v\n", solutions[0])
	fmt.Printf("Looking for %d solutions, found %d solutions.\n", maxSolutionSize, len(solutions)-1)
	for i := 1; i < len(solutions); i++ {
		printSudokuByRaw(squareLength, solutions[i])
	}
}

func Test_SolvePuzzleByRaw(t *testing.T) {
	squareLength := 3
	maxSolutionSize := 5
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

	SolvePuzzleByRaw_Test(squareLength, raw, maxSolutionSize)
}

func Test_solvePuzzle(t *testing.T) {
	squareLength := 3
	var raw string
	var solutions []string
	p := newPuzzle(squareLength)

	raw =
		"..3456789" +
			"456789123" +
			"789123456" +
			"..4365897" +
			"365897214" +
			"897214365" +
			"531642978" +
			"642978531" +
			"978531642" // 2 solutions puzzle

	p.init(raw2digits(raw))
	solutions = strings.Split(p.solvePuzzle(2), string(SOLUTION_PREFIX))
	fmt.Printf("%v\n", solutions[0])
	fmt.Printf("Looking for %d solutions, found %d solutions.\n", 2, len(solutions)-1)
	for i := 1; i < len(solutions); i++ {
		printSudokuByRaw(squareLength, solutions[i])
	}

	raw =
		"....7.94." +
			".7..9...5" +
			"3....5.7." +
			"..74..1.." +
			"463.8...." +
			".....7.8." +
			"8..7....." +
			"7......28" +
			".5..68..." // 188 solutions puzzle

	p.init(raw2digits(raw))
	solutions = strings.Split(p.solvePuzzle(1), string(SOLUTION_PREFIX))
	fmt.Printf("%v\n", solutions[0])
	fmt.Printf("Looking for %d solutions, found %d solutions.\n", 1, len(solutions)-1)
	for i := 1; i < len(solutions); i++ {
		printSudokuByRaw(squareLength, solutions[i])
	}
}

func Test_hasUniqueSolution(t *testing.T) {
	squareLength := 3
	var raw string
	p := newPuzzle(squareLength)

	raw =
		"..3456789" +
			"456789123" +
			"789123456" +
			"..4365897" +
			"365897214" +
			"897214365" +
			"531642978" +
			"642978531" +
			"978531642" // 2 solutions puzzle

	p.init(raw2digits(raw))
	if p.hasUniqueSolution() == true {
		t.Error("2 solutions puzzle, but hasUniqueSolution returns true")
	}

	raw =
		"....7.94." +
			".7..9...5" +
			"3....5.7." +
			"..74..1.." +
			"463.8...." +
			".....7.8." +
			"8..7....." +
			"7......28" +
			".5..68..." // 188 solutions puzzle

	p.init(raw2digits(raw))
	if p.hasUniqueSolution() {
		t.Error("188 solutions puzzle, but hasUniqueSolution returns true")
	}

	raw =
		"......123" +
			"..9......" +
			".....9..." +
			"........." +
			"........." +
			"........." +
			"........." +
			"........." +
			"........." // 0 solutions puzzle

	p.init(raw2digits(raw))
	if p.hasUniqueSolution() == true {
		t.Error("0 solutions puzzle, but hasUniqueSolution returns true")
	}

	raw =
		"........." +
			"..41.26.." +
			".3..5..2." +
			".2..1..3." +
			"..65.41.." +
			".8..7..4." +
			".7..2..6." +
			"..14.35.." +
			"........." // 1 solutions puzzle
	p.init(raw2digits(raw))
	if p.hasUniqueSolution() == false {
		t.Error("1 solutions puzzle, but hasUniqueSolution returns false")
	}
}
