package dlx

/*

                 |           |           |           |
- columns 0 - columns 1 - columns 2 - ......... - columns i -
                 |           |           |           |
            -    x(x0)  -    x      -    x      -    x      -
                 |           |           |           |
                                    -    x(x0)  -    x      -
                                         |           |
                        -    x(x0)  -    x      -    x      -
                             |           |           |
                        -    x(x0)  -    x      -    x      -
                             |           |           |
*/

// data object
type x struct {
	c  *column // column of the x
	x0 *x      // first x in the row
	u  *x      // up x
	d  *x      // down x
	l  *x      // left x
	r  *x      // right x
}

// column object
type column struct {
	x
	i int // index
	n int // x size in the column
}

// an object to hold the matrix and solution
type dlx struct {
	columns []column // columns
	o       []*x     // solution
}

func newDlx(size int) *dlx {
	if (size <= 0) {
		return nil
	}

	columns := make([]column, size + 1) // use column 0 as head
	d := &dlx{columns: columns}

	// init head
	h := &columns[0]
	h.c = h
	// column[0].left = last column
	h.l = &columns[size].x
	// last column.right = column[0]
	columns[size].r = &h.x

	// init columns, but no rows
	for i := 1; i <= size; i ++ {
		c := &columns[i]
		// column.column = column itself
		c.c = c
		// column.index
		c.i = i
		cx := &c.x
		// column.up & column.down = column itself
		c.u, c.d = cx, cx
		// column.left = prev column
		c.l = &h.x
		// prev column.right = curr column
		h.r = cx
		// prev column = curr column
		h = c
	}
	return d
}

// cis: column indexes. Since column[0] is the head, value in the cis should not include 0
func (d *dlx) addRow(cis []int) {
	length := len(cis)
	if length == 0 {
		return
	}

	xs := make([]x, length)
	row := &xs[0]

	for i, ci := range cis {
		c := &d.columns[ci]
		// increase curr.column.n
		c.n++
		x := &xs[i]
		// x.column
		x.c = c
		// x.up = the last x in the column
		x.u = c.u
		// x.down = the column
		x.d = &c.x
		// x.left = prev x in the xs
		// x.right = next x in the xs
		x.l, x.r = &xs[getPrevIndex(i, length)], &xs[getNextIndex(i, length)]
		// x.up.down & x.down.up & x.right.left & x.left.right = x itself
		x.u.d, x.d.u, x.r.l, x.l.r = x, x, x, x
		// reference to first x of the raw
		x.x0 = row
	}
}

func getPrevIndex(i int, length int) int {
	ret := i - 1
	if ret >= 0 {
		return ret
	} else {
		return length - 1
	}
}

func getNextIndex(i int, length int) int {
	ret := i + 1
	if ret < length {
		return ret
	} else {
		return 0
	}
}

func cover(c *column) {
	// column.right.left = column.left
	// column.left.right = column.right
	c.r.l, c.l.r = c.l, c.r

	for i := c.d; i != &c.x; i = i.d {
		for j := i.r; j != i; j = j.r {
			// x.down.up = x.up
			// x.up.down = x.down
			j.d.u, j.u.d = j.u, j.d
			// reduce curr.column.n
			j.c.n--
		}
	}
}

func uncover(c *column) {
	for i := c.u; i != &c.x; i = i.u {
		for j := i.l; j != i; j = j.l {
			// x.down.up & x.up.down = x itself
			j.d.u, j.u.d = j, j
			// increase curr.column.n
			j.c.n++
		}
	}

	c.r.l, c.l.r = &c.x, &c.x
}

// the dlx algorithm
// f(): whether or not stop searching after found a solution. true: stop searching
// If f() return false, it will abandon this solution and continue to search next solution. Cache solution data in f() if necessary
func (d *dlx) search(f func(o []*x) bool) bool {
	h := &d.columns[0]
	hrc := h.r.c
	if hrc == h {
		return f(d.o)
	}

	c := hrc
	min := c.n
	// find the column has minimum x size, it improves overall performance by compare with linear iterator
	for {
		hrc = hrc.r.c
		if hrc == h {
			break
		}
		if hrc.n < min {
			c, min = hrc, hrc.n
		}
	}

	cover(c)
	d.o = append(d.o, nil) // expend d.o length at the end
	length := len(d.o)
	for r := c.d; r != &c.x; r = r.d {
		// set the new item at the end of d.o
		d.o[length - 1] = r
		for j := r.r; j != r; j = j.r {
			cover(j.c)
		}
		if d.search(f) {
			return true
		}
		r = d.o[length - 1]
		c = r.c
		for j := r.l; j != r; j = j.l {
			uncover(j.c)
		}
	}
	d.o = d.o[:length - 1] // remove last item from d.o
	uncover(c)
	return false
}
