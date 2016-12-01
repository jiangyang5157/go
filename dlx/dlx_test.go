package dlx

import (
	"testing"
	"fmt"
)

func Test_NewDlx(t *testing.T) {
	raw := "0123456789" +
		"111111111" +
		"222222222" +
		"333333333" +
		"444444444" +
		"555555555" +
		"666666666" +
		"777777777" +
		"888888888"
	printRaw("#### TEST", raw);
	fmt.Println("\n####")

	d := newDlx(5)
	columns := d.columns;
	for i := range columns {
		fmt.Print(columns[i].i)
	}
	fmt.Println("\n####")
	h := &d.columns[0]
	c := h.r.c
	for ;c.x != h.x; c = c.r.c {
		fmt.Print(c.i)
	}
	fmt.Println("\n####")
	fmt.Printf("%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v\n", int('0'), int('1'), int('2'), int('3'), int('4'), int('5'), int('6'), int('7'), int('8'), int('9'), int('.'))
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