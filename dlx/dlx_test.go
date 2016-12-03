package dlx

import (
	"testing"
	"fmt"
)

func Test_Ascii(t *testing.T) {
	//0,  1,  2,  3,  4,  5,  6,  7,  8,  9,  .
	//48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 46
	fmt.Printf("%v,  %v,  %v,  %v,  %v,  %v,  %v,  %v,  %v,  %v,  %v\n", 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, ".")
	fmt.Printf("%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v\n", int('0'), int('1'), int('2'), int('3'), int('4'), int('5'), int('6'), int('7'), int('8'), int('9'), int('.'))
}

func Test_newDlx(t *testing.T) {
	d := newDlx(5)
	columns := d.columns;
	for i := range columns {
		fmt.Print(columns[i].i)
	}
	fmt.Println("")
	h := &d.columns[0]
	c := h.r.c
	for ; c.x != h.x; c = c.r.c {
		fmt.Print(c.i)
	}
	fmt.Println("")
}