package dlx

/*
Constraints example: 9x9 puzzle (squareLength = 3)
1. Each cell must has a digit: 9 * 9 = 81 constraints in column 1-81
2. Each row must has [1, 9]: 9 * 9 = 81 constraints in column 82-162
3. Each column must has [1, 9]: 9 * 9 = 81 constraints in column 163-243
4. Each square must has [1, 9]: 9 * 9 = 81 constraints in column 244-324
*/
type puzzle struct {
	dlx
	squareLength int // > 0
	edgeLength   int // squareLength * squareLength
	cellSize     int // edgeLength * edgeLength
	offset1      int // 0
	offset2      int // offset1 + cellSize
	offset3      int // offset2 + cellSize
	offset4      int // offset3 + cellSize
}

func newPuzzle(squareLength int) *puzzle {
	edgeLength := squareLength * squareLength
	cellSize := edgeLength * edgeLength
	return &puzzle{
		squareLength: squareLength,
		edgeLength:   edgeLength,
		cellSize:     cellSize,
		offset1:      cellSize * 0,
		offset2:      cellSize * 1,
		offset3:      cellSize * 2,
		offset4:      cellSize * 3,
	}
}

func (p *puzzle) init(digits *[]int) {
	columnSize := p.offset4 + p.cellSize
	p.dlx = *newDlx(columnSize)
	for r, i := 0, 0; r < p.edgeLength; r++ {
		for c := 0; c < p.edgeLength; c, i = c + 1, i + 1 {
			s := p.getSquareIndex(r, c)
			digit := (*digits)[i]
			p.addDigit(digit, i, r, c, s)
		}
	}
}

func (p *puzzle) addDigit(digit int, i int, r int, c int, s int) {
	if digit >= 1 && digit <= p.edgeLength {
		// valid digit
		p.addRow([]int{
			p.offset1 + i + 1,
			p.offset2 + r * p.edgeLength + digit,
			p.offset3 + c * p.edgeLength + digit,
			p.offset4 + s * p.edgeLength + digit})
	} else {
		// unknown digit, consider all possibilities
		for digit = 1; digit <= p.edgeLength; digit++ {
			p.addRow([]int{
				p.offset1 + i + 1,
				p.offset2 + r * p.edgeLength + digit,
				p.offset3 + c * p.edgeLength + digit,
				p.offset4 + s * p.edgeLength + digit})
		}
	}
}

func (p *puzzle) getSquareIndex(r int, c int) int {
	return r / p.squareLength * p.squareLength + c / p.squareLength
}

func (p *puzzle) getCellIndex(r int, c int) int {
	return r * p.edgeLength + c
}

func (p *puzzle) resetSolution() {
	p.o = p.o[:0]
}

func (p *puzzle) hasUniqueSolution() bool {
	solutionCount := 0
	p.search(func(o []*x) bool {
		solutionCount++
		return solutionCount > 1
	})

	return solutionCount == 1
}

func raw2digits(raw *string) *[]int {
	rawLength := len(*raw)
	var digits []int = make([]int, rawLength)
	for i := 0; i < rawLength; i++ {
		digits[i] = int((*raw)[i] - '0')
	}
	return &digits
}

func digits2raw(digits *[]int) string {
	digitLength := len(*digits)
	var bs []byte = make([]byte, digitLength)
    	for i := 0; i < digitLength; i++ {
        	bs[i] = byte((*digits)[i])
    	}
	fmt.Println("bs", bs)
	ret := string(bs)
	fmt.Println("ret", ret)
	return ret
}
