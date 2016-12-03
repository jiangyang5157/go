package dlx

import (
	"testing"
	"fmt"
	"strings"
)

func Test_Ascii(t *testing.T) {
	//0,  1,  2,  3,  4,  5,  6,  7,  8,  9,  .
	//48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 46
	fmt.Printf("%v,  %v,  %v,  %v,  %v,  %v,  %v,  %v,  %v,  %v,  %v\n", 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, ".")
	fmt.Printf("%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v\n", int('0'), int('1'), int('2'), int('3'), int('4'), int('5'), int('6'), int('7'), int('8'), int('9'), int('.'))
}

func printSudoku(squareLength int, raw string) {
	fmt.Println("================================================================")
	switch squareLength {
	case 1:
		for r, i := 0, 0; r < 1; r, i = r + 1, i + 1 {
			fmt.Printf("%c\n", raw[i])
		}
	case 2:
		for r, i := 0, 0; r < 4; r, i = r + 1, i + 4 {
			fmt.Printf("%c %c | %c %c\n",
				raw[i], raw[i + 1],
				raw[i + 2], raw[i + 3])
			if r == 1 {
				fmt.Println("----+----")
			}
		}
	case 3:
		for r, i := 0, 0; r < 9; r, i = r + 1, i + 9 {
			fmt.Printf("%c %c %c | %c %c %c | %c %c %c\n",
				raw[i], raw[i + 1], raw[i + 2],
				raw[i + 3], raw[i + 4], raw[i + 5],
				raw[i + 6], raw[i + 7], raw[i + 8])
			if r == 2 || r == 5 {
				fmt.Println("------+-------+------")
			}
		}
	case 4:
		// Ascii for chars after '9' can convert into integer by minus '0' for clear display.
		for r, i := 0, 0; r < 16; r, i = r + 1, i + 16 {
			fmt.Printf("%c %c %c %c | %c %c %c %c | %c %c %c %c | %c %c %c %c\n",
				raw[i], raw[i + 1], raw[i + 2], raw[i + 3],
				raw[i + 4], raw[i + 5], raw[i + 6], raw[i + 7],
				raw[i + 8], raw[i + 9], raw[i + 10], raw[i + 11],
				raw[i + 12], raw[i + 13], raw[i + 14], raw[i + 15])
			if r == 3 || r == 7 || r == 11 {
				fmt.Println("--------+---------+---------+--------")
			}
		}
	default:
		fmt.Printf("squareLength: %v\nraw: %v\n", squareLength, raw)
	}
}

func solveSudukuTest(squareLength int, raw string, solutionSize int) {
	printSudoku(squareLength, raw)
	solutions := strings.Split(SolveSudoku(squareLength, raw, solutionSize), string(PREFIX_SOLUTION))
	length := len(solutions)
	// [0] contains massage
	fmt.Printf("%v\n", solutions[0])
	fmt.Printf("Looking for %d solutions, found %d solutions.\n", solutionSize, length - 1)
	for i := 1; i < length; i++ {
		printSudoku(squareLength, solutions[i])
	}
}

func Test_solve(t *testing.T) {
	squareLength := 3
	solutionSize := 2
	raw :=
	//"......123" +
	//"..9......" +
	//".....9..." +
	//"........." +
	//"........." +
	//"........." +
	//"........." +
	//"........." +
	//"........." // 0 solutions Sudoku Puzzle

	//"........." +
	//"..41.26.." +
	//".3..5..2." +
	//".2..1..3." +
	//"..65.41.." +
	//".8..7..4." +
	//".7..2..6." +
	//"..14.35.." +
	//"........." // 1 solutions Sudoku Puzzle

	"..3456789" +
	"456789123" +
	"789123456" +
	"..4365897" +
	"365897214" +
	"897214365" +
	"531642978" +
	"642978531" +
	"978531642" // 2 solutions Sudoku Puzzle

	//"....7.94." +
	//".7..9...5" +
	//"3....5.7." +
	//"..74..1.." +
	//"463.8...." +
	//".....7.8." +
	//"8..7....." +
	//"7......28" +
	//".5..68..." // 188 solutions Sudoku puzzle

	//"." // 1 solutions Sudoku puzzle

	//"...." +
	//".4.." +
	//"2..." +
	//"..43" // 0 solutions Sudoku puzzle

	//"...." +
	//".4.." +
	//"2..." +
	//"...3" // 3 solutions Sudoku puzzle

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
	//"................" // ? solutions Sudoku puzzle

	solveSudukuTest(squareLength, raw, solutionSize)
}
