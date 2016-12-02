package dlx

import (
	"testing"
	"fmt"
)

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

func printRawSudoku3(title, raw string) {
	fmt.Println(title)
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

func Test_stuff(t *testing.T) {
	raw := ".......12........3..23..4....18....5.6..7.8.......9.....85.....9...4.5..47...6..."
	printRawSudoku3("#### test printRawSudoku3", raw)

	fmt.Println("#### ascii")
	//0,  1,  2,  3,  4,  5,  6,  7,  8,  9,  .
	//48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 46
	fmt.Printf("%v,  %v,  %v,  %v,  %v,  %v,  %v,  %v,  %v,  %v,  %v\n", 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, ".")
	fmt.Printf("%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v\n", int('0'), int('1'), int('2'), int('3'), int('4'), int('5'), int('6'), int('7'), int('8'), int('9'), int('.'))
}

func Test_sudoku(t *testing.T) {
	raw := ".......12........3..23..4....18....5.6..7.8.......9.....85.....9...4.5..47...6..."
	printRawSudoku3("#### raw", raw)
	solution := solve(3, raw)
	printRawSudoku3("#### solution", solution)
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

	d.search()
	bs := make([]byte, len(d.o))
	for _, o := range d.o {
		x0 := o.x0
		x0ci := x0.c.i // [offsetConstraint1 + 1, offsetConstraint2] cell constraints - index for byte
		x0rci := x0.r.c.i // [offsetConstraint2 + 1, offsetConstraint3] row constraints - append by raw
		bs[x0ci - 1] = byte(x0rci % edgeLength) + '0'
	}
	return string(bs)
}
