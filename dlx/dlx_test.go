package dlx

import (
	"testing"
	"fmt"
)

func Test_stuff(t *testing.T) {
	testRaw := ".......12........3..23..4....18....5.6..7.8.......9.....85.....9...4.5..47...6..."
	printRaw("#### test printRaw", testRaw)

	fmt.Println("#### test values of char")
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