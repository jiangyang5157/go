package dlx

/*
Constraints example: 9x9 Sudoku (squareLength = 3)
1. Each cell must has a digit: 9 * 9 = 81 constraints in column 1-81
2. Each row must has [1, 9]: 9 * 9 = 81 constraints in column 82-162
3. Each column must has [1, 9]: 9 * 9 = 81 constraints in column 163-243
4. Each square must has [1, 9]: 9 * 9 = 81 constraints in column 244-324
 */
type SudokuDlx struct {
	dlx
	squareLength int // > 0
	edgeLength   int // squareLength * squareLength
	cellSize     int // edgeLength * edgeLength
	offset1      int // 0
	offset2      int // offset1 + cellSize
	offset3      int // offset2 + cellSize
	offset4      int // offset3 + cellSize
}

func newSudokuDlx(squareLength int) *SudokuDlx {
	edgeLength := squareLength * squareLength
	cellSize := edgeLength * edgeLength
	sd := &SudokuDlx{
		squareLength: squareLength,
		edgeLength: edgeLength,
		cellSize: cellSize,
		offset1: cellSize * 0,
		offset2: cellSize * 1,
		offset3: cellSize * 2,
		offset4: cellSize * 3,
	}
	return sd
}

func (sd *SudokuDlx) initializeDlx(raw string) {
	columnSize := sd.offset4 + sd.cellSize
	sd.dlx = *newDlx(columnSize)
	for r, i := 0, 0; r < sd.edgeLength; r++ {
		for c := 0; c < sd.edgeLength; c, i = c + 1, i + 1 {
			s := sd.squareIndex(r, c)
			digit := int(raw[i] - '0')
			sd.addDigit(digit, i, r, c, s)
		}
	}
}

func (sd *SudokuDlx)addDigit(digit int, i int, r int, c int, s int) {
	if digit >= 1 && digit <= sd.edgeLength {
		// valid digit
		sd.addRow([]int{
			sd.offset1 + i + 1,
			sd.offset2 + r * sd.edgeLength + digit,
			sd.offset3 + c * sd.edgeLength + digit,
			sd.offset4 + s * sd.edgeLength + digit})
	} else {
		// unknown digit, consider all possibilities
		for digit = 1; digit <= sd.edgeLength; digit++ {
			sd.addRow([]int{
				sd.offset1 + i + 1,
				sd.offset2 + r * sd.edgeLength + digit,
				sd.offset3 + c * sd.edgeLength + digit,
				sd.offset4 + s * sd.edgeLength + digit})
		}
	}
}

func (sd *SudokuDlx) resetSolution() {
	sd.o = sd.o[:0]
}

func (sd *SudokuDlx) squareIndex(r int, c int) int {
	return r / sd.squareLength * sd.squareLength + c / sd.squareLength
}
