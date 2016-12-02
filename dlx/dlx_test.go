package dlx

import (
	"testing"
	"fmt"
	"strings"
)

func Test_stuff(t *testing.T) {
	raw := ".......12........3..23..4....18....5.6..7.8.......9.....85.....9...4.5..47...6..."
	fmt.Println("#### printSudoku")
	printSudoku3(raw)

	fmt.Println("#### ascii")
	//0,  1,  2,  3,  4,  5,  6,  7,  8,  9,  .
	//48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 46
	fmt.Printf("%v,  %v,  %v,  %v,  %v,  %v,  %v,  %v,  %v,  %v,  %v\n", 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, ".")
	fmt.Printf("%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v\n", int('0'), int('1'), int('2'), int('3'), int('4'), int('5'), int('6'), int('7'), int('8'), int('9'), int('.'))
}

func Test_NewDlx(t *testing.T) {
	fmt.Println("#### test newDlx")
	d := newDlx(5)
	columns := d.columns;
	for i := range columns {
		fmt.Print(columns[i].i)
	}
	fmt.Println("\n####")
	h := &d.columns[0]
	c := h.r.c
	for ; c.x != h.x; c = c.r.c {
		fmt.Print(c.i)
	}
	fmt.Println("\n####")
}

func printSudoku1(raw string) {
	for r, i := 0, 0; r < 1; r, i = r + 1, i + 1 {
		fmt.Printf("%c\n", raw[i])
	}
}

func printSudoku2(raw string) {
	for r, i := 0, 0; r < 4; r, i = r + 1, i + 4 {
		fmt.Printf("%c %c | %c %c\n",
			raw[i], raw[i + 1],
			raw[i + 2], raw[i + 3])
		if r == 1 {
			fmt.Println("----+----")
		}
	}
}

func printSudoku3(raw string) {
	for r, i := 0, 0; r < 9; r, i = r + 1, i + 9 {
		fmt.Printf("%c %c %c | %c %c %c | %c %c %c\n",
			raw[i], raw[i + 1], raw[i + 2],
			raw[i + 3], raw[i + 4], raw[i + 5],
			raw[i + 6], raw[i + 7], raw[i + 8])
		if r == 2 || r == 5 {
			fmt.Println("------+-------+------")
		}
	}
}

// the ascii for chars after '9' should convert into integer by minus '0' for display.
func printSudoku4(raw string) {
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
}

func Test_sudoku4(t *testing.T) {
	raw :=
		"................" +
			"................" +
			"................" +
			"................" +
			"................" +
			"................" +
			"................" +
			"................" +
			"................" +
			"................" +
			"................" +
			"................" +
			"................" +
			"................" +
			"................" +
			"................"
	fmt.Println("#### raw")
	printSudoku4(raw)
	solutions := strings.Split(solve(4, raw), ",")
	length := len(solutions)
	if length >= 2 {
		for i := 0; i < length - 1; i++ {
			fmt.Println("#### solution")
			printSudoku4(solutions[i])
		}
	}
}

func Test_sudoku3(t *testing.T) {
	raw :=
		"....7.94." +
			".7..9...5" +
			"3....5.7." +
			"..74..1.." +
			"463.8...." +
			".....7.8." +
			"8..7....." +
			"7......28" +
			".5..68..."
	fmt.Println("#### raw")
	printSudoku3(raw)
	solutions := strings.Split(solve(3, raw), ",")
	length := len(solutions)
	if length >= 2 {
		for i := 0; i < length - 1; i++ {
			fmt.Println("#### solution")
			printSudoku3(solutions[i])
		}
	}
}

func Test_sudoku3withZeroSolution(t *testing.T) {
	raw :=
		"......123" +
			"..9......" +
			".....9..." +
			"........." +
			"........." +
			"........." +
			"........." +
			"........." +
			"........."
	fmt.Println("#### raw")
	printSudoku3(raw)
	solutions := strings.Split(solve(3, raw), ",")
	length := len(solutions)
	if length >= 2 {
		for i := 0; i < length - 1; i++ {
			fmt.Println("#### solution")
			printSudoku3(solutions[i])
		}
	}
}

func Test_sudoku2(t *testing.T) {
	raw :=
		"...." +
			".4.." +
			"2..." +
			"...3"
	fmt.Println("#### raw")
	printSudoku2(raw)
	solutions := strings.Split(solve(2, raw), ",")
	length := len(solutions)
	if length >= 2 {
		for i := 0; i < length - 1; i++ {
			fmt.Println("#### solution")
			printSudoku2(solutions[i])
		}
	}
}

func Test_sudoku2WithZeroSolution(t *testing.T) {
	raw :=
		"...." +
			".4.." +
			"2..." +
			"..43"
	fmt.Println("#### raw")
	printSudoku2(raw)
	solutions := strings.Split(solve(2, raw), ",")
	length := len(solutions)
	if length >= 2 {
		for i := 0; i < length - 1; i++ {
			fmt.Println("#### solution")
			printSudoku2(solutions[i])
		}
	}
}

func Test_sudoku1(t *testing.T) {
	raw := "."
	fmt.Println("#### raw")
	printSudoku1(raw)
	solutions := strings.Split(solve(1, raw), ",")
	length := len(solutions)
	if length >= 2 {
		for i := 0; i < length - 1; i++ {
			fmt.Println("#### solution")
			printSudoku1(solutions[i])
		}
	}
}

/*
Eg: 9x9 sudoku (s = 3)
1. Each cell must has a digit: 9 * 9 = 81 constraints in column 1-81
2. Each row must has [1, 9]: 9 * 9 = 81 constraints in column 82-162
3. Each column must has [1, 9]: 9 * 9 = 81 constraints in column 163-243
4. Each square must has [1, 9]: 9 * 9 = 81 constraints in column 244-324
*/
func solve(s int, raw string) string {
	if (s <= 0) {
		return "error: s"
	}

	edgeLength := s * s
	cellSize := edgeLength * edgeLength
	if (len(raw) != cellSize) {
		return "error: raw"
	}

	offsetConstraint1 := 0
	offsetConstraint2 := offsetConstraint1 + cellSize
	offsetConstraint3 := offsetConstraint2 + cellSize
	offsetConstraint4 := offsetConstraint3 + cellSize
	columnSize := offsetConstraint4 + cellSize
	d := newDlx(columnSize)

	for r, i := 0, 0; r < edgeLength; r++ {
		// r = [0, edgeLength - 1]
		// i = [0, cellSize - 1]
		for c := 0; c < edgeLength; c, i = c + 1, i + 1 {
			// c: [0, edgeLength - 1]
			// square: [0, edgeLength - 1]
			square := r / s * s + c / s

			digit := int(raw[i] - '0')
			if digit >= 1 && digit <= edgeLength {
				d.addRow([]int{
					offsetConstraint1 + i + 1,
					offsetConstraint2 + r * edgeLength + digit,
					offsetConstraint3 + c * edgeLength + digit,
					offsetConstraint4 + square * edgeLength + digit})
			} else {
				// consider all possibilities
				for digit = 1; digit <= edgeLength; digit++ {
					d.addRow([]int{
						offsetConstraint1 + i + 1,
						offsetConstraint2 + r * edgeLength + digit,
						offsetConstraint3 + c * edgeLength + digit,
						offsetConstraint4 + square * edgeLength + digit})
				}
			}
		}
	}

	var buffer []byte
	count := 0
	need := 1
	ok := d.search(func(o []*x) bool {
		fmt.Println("FOUND")
		count++

		b := make([]byte, len(o))
		for _, o := range d.o {
			x0 := o.x0
			x0ci := x0.c.i // [offsetConstraint1 + 1, offsetConstraint2] cell constraints - index for byte
			x0rci := x0.r.c.i // [offsetConstraint2 + 1, offsetConstraint3] row constraints - append by raw
			b[x0ci - 1] = byte((x0rci - 1) % edgeLength) + '1'
		}
		buffer = append(buffer, b...)
		buffer = append(buffer, ',')

		return count >= need
	})
	fmt.Println(ok)
	if (!ok) {
		return ""
	}
	return string(buffer)
}

func Test_sudoku(t *testing.T) {
	raw :=
	// 188 sulutions
		"....7.94." +
			".7..9...5" +
			"3....5.7." +
			"..74..1.." +
			"463.8...." +
			".....7.8." +
			"8..7....." +
			"7......28" +
			".5..68..."
	// 2 sulutions
	//"..3456789" +
	//"456789123" +
	//"789123456" +
	//"..4365897" +
	//"365897214" +
	//"897214365" +
	//"531642978" +
	//"642978531" +
	//"978531642"
	// 1 sulution
	//"........." +
	//"..41.26.." +
	//".3..5..2." +
	//".2..1..3." +
	//"..65.41.." +
	//".8..7..4." +
	//".7..2..6." +
	//"..14.35.." +
	//"........."
	fmt.Println("#### raw")
	printSudoku3(raw)
	solutions := strings.Split(solve(3, raw), ",")
	length := len(solutions)
	if length >= 2 {
		for i := 0; i < length - 1; i++ {
			fmt.Println("#### solution")
			printSudoku3(solutions[i])
		}
	}
}

// Eg: 9x9 sudoku
// 0 sulution
//"......123" +
//"..9......" +
//".....9..." +
//"........." +
//"........." +
//"........." +
//"........." +
//"........." +
//"........."
//
// 1 sulution
//"........." +
//"..41.26.." +
//".3..5..2." +
//".2..1..3." +
//"..65.41.." +
//".8..7..4." +
//".7..2..6." +
//"..14.35.." +
//"........."
//
// 2 sulutions
//"..3456789" +
//"456789123" +
//"789123456" +
//"..4365897" +
//"365897214" +
//"897214365" +
//"531642978" +
//"642978531" +
//"978531642"
//
// 188 sulutions
//"....7.94." +
//".7..9...5" +
//"3....5.7." +
//"..74..1.." +
//"463.8...." +
//".....7.8." +
//"8..7....." +
//"7......28" +
//".5..68..."
