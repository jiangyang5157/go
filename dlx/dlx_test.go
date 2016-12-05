package dlx

import (
	"fmt"
	"testing"
)

func Test_newDlx(t *testing.T) {
	d := newDlx(5)
	columns := d.columns
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
