package dlx

// prefix < '0' && prefix != whatever representing unknown digit in the raw
const SOLUTION_PREFIX byte = '#'

func SolveSudoku(squareLength int, raw string, solutionSize int) string {
	rawLength := len(raw)
	var digits []int = make([]int, rawLength)
	for i := 0; i < rawLength; i++ {
		digits[i] = int(raw[i] - '0')
	}
	return solveSudoku(squareLength, digits, solutionSize)
}

func solveSudoku(squareLength int, digits []int, solutionSize int) string {
	if (solutionSize < 1) {
		return "No action required"
	}
	if (squareLength < 1) {
		return "Invalid Sudoku puzzle"
	}

	p := newSudoku(squareLength)
	if (len(digits) != p.cellSize) {
		return "Invalid Sudoku data"
	}
	p.init(digits)

	var ret []byte
	solutionCount := 0
	p.search(func(o []*x) bool {
		bs := make([]byte, len(o)) // o.len = cellSize
		for _, o := range p.o {
			x0 := o.x0
			x0ci := x0.c.i // x0ci = [offset1 + 1, offset2]
			x0rci := x0.r.c.i // x0rci = [offset2 + 1, offset3]
			// bytes append by raw, index = [0, cellSize - 1]
			bs[x0ci - 1] = byte((x0rci - 1) % p.edgeLength) + '1'
		}
		ret = append(ret, SOLUTION_PREFIX)
		ret = append(ret, bs...)
		solutionCount++
		return solutionCount >= solutionSize
	})
	return string(ret)
}

func (p *puzzle) uniqueSolution(digits []int) bool {
	solutionCount := 0
	p.search(func(o []*x) bool {
		solutionCount++
		return solutionCount > 1
	})
	return solutionCount == 1
}

func (p *puzzle) resetSolution() {
	p.o = p.o[:0]
}