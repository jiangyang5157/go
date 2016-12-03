package dlx

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
	edgeLength := squareLength * squareLength
	cellSize := edgeLength * edgeLength
	if (len(raw) != cellSize) {
		return "Invalid Sudoku raw"
	}

	d := newSudokuDlx(raw, squareLength, edgeLength, cellSize)
	var ret []byte
	solutionCount := 0
	d.search(func(o []*x) bool {
		bs := make([]byte, len(o)) // o.len = cellSize
		for _, o := range d.o {
			x0 := o.x0
			x0ci := x0.c.i // x0ci = [offset1 + 1, offset2]
			x0rci := x0.r.c.i // x0rci = [offset2 + 1, offset3]
			// bytes append by raw, index = [0, cellSize - 1]
			bs[x0ci - 1] = byte((x0rci - 1) % edgeLength) + '1'
		}
		ret = append(ret, PUZZLE_PREFIX)
		ret = append(ret, bs...)
		solutionCount++
		return solutionCount >= solutionSize
	})
	return string(ret)
}