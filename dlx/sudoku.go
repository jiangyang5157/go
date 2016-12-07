package dlx

import (
	"time"
	"math/rand"
)

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
	offset1 := cellSize * 0
	offset2 := cellSize * 1
	offset3 := cellSize * 2
	offset4 := cellSize * 3
	p := &puzzle{
		squareLength: squareLength,
		edgeLength:   edgeLength,
		cellSize:     cellSize,
		offset1:      offset1,
		offset2:      offset2,
		offset3:      offset3,
		offset4:      offset4,
	}
	return p
}

func (p *puzzle) init(digits []int) {
	columnSize := p.offset4 + p.cellSize
	p.dlx = *newDlx(columnSize)

	edgeLength := p.edgeLength
	for r, i := 0, 0; r < edgeLength; r++ {
		for c := 0; c < edgeLength; c, i = c + 1, i + 1 {
			s := p.squareIndex(r, c)
			digit := digits[i]
			p.addDigit(digit, i, r, c, s)
		}
	}
}

func (p *puzzle) addDigit(digit int, i int, r int, c int, s int) {
	edgeLength := p.edgeLength
	if digit >= 1 && digit <= edgeLength {
		// valid digit
		p.addRow([]int{
			p.offset1 + i + 1,
			p.offset2 + r * edgeLength + digit,
			p.offset3 + c * edgeLength + digit,
			p.offset4 + s * edgeLength + digit})
	} else {
		// unknown digit, consider all possibilities
		for digit = 1; digit <= edgeLength; digit++ {
			p.addRow([]int{
				p.offset1 + i + 1,
				p.offset2 + r * edgeLength + digit,
				p.offset3 + c * edgeLength + digit,
				p.offset4 + s * edgeLength + digit})
		}
	}
}

func (p *puzzle) squareIndex(r int, c int) int {
	squareLength := p.squareLength
	return r / squareLength * squareLength + c / squareLength
}

func (p *puzzle) cellIndex(r int, c int) int {
	return r * p.edgeLength + c
}

func (p *puzzle) rcIndex(cellIndex int) (int, int) {
	edgeLength := p.edgeLength
	return cellIndex / edgeLength, cellIndex % edgeLength
}

func raw2digits(raw string) []int {
	length := len(raw)
	var digits []int = make([]int, length)
	for i := 0; i < length; i++ {
		digits[i] = int(raw[i] - '0')
	}
	return digits
}

func digits2raw(digits []int) string {
	length := len(digits)
	var bs []byte = make([]byte, length)
	for i := 0; i < length; i++ {
		bs[i] = byte(digits[i]) + '0'
	}
	return string(bs)
}

func disorderArray(array []int) []int {
	rand.Seed(time.Now().Unix())
	length := len(array)
	for i := 0; i < length; i++ {
		random := rand.Intn(length)
		array[i], array[random] = array[random], array[i]
	}
	return array
}