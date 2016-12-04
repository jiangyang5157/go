package dlx

/*
Constraints example: 9x9 Sudoku (squareLength = 3)
1. Each cell must has a digit: 9 * 9 = 81 constraints in column 1-81
2. Each row must has [1, 9]: 9 * 9 = 81 constraints in column 82-162
3. Each column must has [1, 9]: 9 * 9 = 81 constraints in column 163-243
4. Each square must has [1, 9]: 9 * 9 = 81 constraints in column 244-324

squareLength > 0
edgeLength = squareLength * squareLength
cellSize = edgeLength * edgeLength
i = [0, cellSize - 1]
r = [0, edgeLength - 1]
c = [0, edgeLength - 1]
s = [0, edgeLength - 1]
offset1 := 0
offset2 := offset1 + cellSize
offset3 := offset2 + cellSize
offset4 := offset3 + cellSize
 */

func newSudokuDlx(raw string, squareLength int, edgeLength int, cellSize int) *dlx {
	offset1 := 0
	offset2 := offset1 + cellSize
	offset3 := offset2 + cellSize
	offset4 := offset3 + cellSize
	columnSize := offset4 + cellSize
	d := newDlx(columnSize)
	for r, i := 0, 0; r < edgeLength; r++ {
		for c := 0; c < edgeLength; c, i = c + 1, i + 1 {
			s := squareIndex(squareLength, r, c)
			digit := int(raw[i] - '0')
			d.addSudokuDigit(digit, i, r, c, s, edgeLength, offset1, offset2, offset3, offset4)
		}
	}
	return d
}

func (d *dlx)addSudokuDigit(digit int, i int, r int, c int, s int, edgeLength int, offset1 int, offset2 int, offset3 int, offset4 int) {
	if digit >= 1 && digit <= edgeLength {
		// valid digit
		d.addRow([]int{
			offset1 + i + 1,
			offset2 + r * edgeLength + digit,
			offset3 + c * edgeLength + digit,
			offset4 + s * edgeLength + digit})
	} else {
		// unknown digit, consider all possibilities
		for digit = 1; digit <= edgeLength; digit++ {
			d.addRow([]int{
				offset1 + i + 1,
				offset2 + r * edgeLength + digit,
				offset3 + c * edgeLength + digit,
				offset4 + s * edgeLength + digit})
		}
	}
}

// reset d.o without wipe out columns and rows structure
func (d *dlx) reset() {
	d.o = d.o[:0]
}

func squareIndex(squareLength int, r int, c int) int {
	return r / squareLength * squareLength + c / squareLength
}
