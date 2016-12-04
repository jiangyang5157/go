package dlx

// prefix <'0' && prefix != whatever representing unknown digit in the raw
const SOLUTION_PREFIX byte = '#'

func SimpleSolveSudoku(squareLength int, raw string) string {
	return SolveSudoku(squareLength, raw, 1)
}

func SolveSudoku(squareLength int, raw string, solutionSize int) string {
	if (solutionSize < 1) {
		return "No action required"
	}
	if (squareLength < 1) {
		return "Invalid Sudoku puzzle"
	}

	sd := newSudokuDlx(squareLength)
	if (len(raw) != sd.cellSize) {
		return "Invalid Sudoku raw"
	}

	sd.initializeDlx(raw)
	var ret []byte
	solutionCount := 0
	sd.search(func(o []*x) bool {
		bs := make([]byte, len(o)) // o.len = cellSize
		for _, o := range sd.o {
			x0 := o.x0
			x0ci := x0.c.i // x0ci = [offset1 + 1, offset2]
			x0rci := x0.r.c.i // x0rci = [offset2 + 1, offset3]
			// bytes append by raw, index = [0, cellSize - 1]
			bs[x0ci - 1] = byte((x0rci - 1) % sd.edgeLength) + '1'
		}
		ret = append(ret, SOLUTION_PREFIX)
		ret = append(ret, bs...)
		solutionCount++
		return solutionCount >= solutionSize
	})
	return string(ret)
}