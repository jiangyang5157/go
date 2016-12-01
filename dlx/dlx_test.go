package dlx

import (
	"testing"
	"fmt"
)

func Test_stuff(t *testing.T) {
	raw := ".......12........3..23..4....18....5.6..7.8.......9.....85.....9...4.5..47...6..."
	printRaw("#### test printRaw", raw)

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

func Test_sudoku(t *testing.T) {
	/*
	eg: 9 x 9 sudoku (Level 3)
        1. Each cell must has a digit: 9 * 9 = 81 constraints in column 1-81
        2. Each row must has [1, 9]: 9 * 9 = 81 constraints in column 82-162
        3. Each column must has [1, 9]: 9 * 9 = 81 constraints in column 163-243
        4. Each square must has [1, 9]: 9 * 9 = 81 constraints in column 244-324
	*/
	raw := ".......12........3..23..4....18....5.6..7.8.......9.....85.....9...4.5..47...6..."
	printRaw("#### raw", raw)
	d := newDlx(324)
	for r, i := 0, 0; r < 9; r++ {
		for c := 0; c < 9; c, i = c + 1, i + 1 {
			// i:[0, 80]
			// r: [0, 8]
			// c: [0, 8]
			// square: [0, 8]
			square := r / 3 * 3 + c / 3
			digit := int(raw[i] - '0')
			if digit >= 1 && digit <= 9 {
				d.addRow([]int{
					i + 1,
					81 + r * 9 + digit,
					162 + c * 9 + digit,
					243 + square * 9 + digit})
			} else {
				// consider all possibilities
				for digit = 1; digit <= 9; digit++ {
					d.addRow([]int{
						i + 1,
						81 + r * 9 + digit,
						162 + c * 9 + digit,
						243 + square * 9 + digit})
				}
			}
		}
	}
	d.search()
	length := len(d.o)
	b := make([]byte, length)
	for _, o := range d.o {
		x0 := o.x0
		x0ci := x0.c.i // [1, 81] cell constraints - index for byte
		x0rci := x0.r.c.i // [82, 162] row constraints - append by raw
		b[x0ci - 1] = byte(x0rci % 9) + '0'
	}
	printRaw("#### solution", string(b))
}

func printRaw(title, raw string) {
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