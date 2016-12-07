package dlx

// prefix < '0' && prefix != whatever representing unknown digit in the raw
const SOLUTION_PREFIX byte = '#'

func SolvePuzzleByRaw(squareLength int, raw string, maxSolutionSize int) string {
	return SolvePuzzleByDigits(squareLength, raw2digits(raw), maxSolutionSize)
}

func SolvePuzzleByDigits(squareLength int, digits []int, maxSolutionSize int) string {
	if maxSolutionSize <= 0 {
		return "No action required"
	}
	if squareLength <= 0 {
		return "Invalid puzzle"
	}

	p := newPuzzle(squareLength)
	if len(digits) != p.cellSize {
		return "Invalid data"
	}
	p.addDigits(digits)
	return p.solvePuzzle(maxSolutionSize)
}

func (p *puzzle) solvePuzzle(maxSolutionSize int) string {
	var ret []byte
	solutionCount := 0
	p.search(func(o []*x) bool {
		bs := make([]byte, p.cellSize)
		for _, o := range p.o {
			x0 := o.x0
			x0ci := x0.c.i    // x0ci = [offset1 + 1, offset2]
			x0rci := x0.r.c.i // x0rci = [offset2 + 1, offset3]
			// bytes append by raw, index = [0, cellSize - 1]
			bs[x0ci - 1] = byte((x0rci - 1) % p.edgeLength) + '1'
		}
		ret = append(ret, SOLUTION_PREFIX)
		ret = append(ret, bs...)
		solutionCount++
		return solutionCount >= maxSolutionSize
	})
	return string(ret)
}

func (p *puzzle) hasUniqueSolution() bool {
	solutionCount := 0
	p.search(func(o []*x) bool {
		solutionCount++
		return solutionCount > 1
	})
	return solutionCount == 1
}