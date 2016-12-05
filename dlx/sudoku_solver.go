package dlx

// prefix < '0' && prefix != whatever representing unknown digit in the raw
const SOLUTION_PREFIX byte = '#'

func SolvePuzzleByRaw(squareLength int, raw string, solutionSize int) string {
	return SolvePuzzleByDigits(squareLength, raw2digits(&raw), solutionSize)
}

func SolvePuzzleByDigits(squareLength int, digits *[]int, solutionSize int) string {
	if solutionSize < 1 {
		return "No action required"
	}
	if squareLength < 1 {
		return "Invalid puzzle"
	}

	p := newPuzzle(squareLength)
	if len(*digits) != p.cellSize {
		return "Invalid data"
	}

	p.init(digits)
	var ret []byte
	solutionCount := 0
	p.search(func(o []*x) bool {
		bs := make([]byte, len(o)) // o.len = cellSize
		for _, o := range p.o {
			x0 := o.x0
			x0ci := x0.c.i    // x0ci = [offset1 + 1, offset2]
			x0rci := x0.r.c.i // x0rci = [offset2 + 1, offset3]
			// bytes append by raw, index = [0, cellSize - 1]
			bs[x0ci-1] = byte((x0rci-1)%p.edgeLength) + '1'
		}
		ret = append(ret, SOLUTION_PREFIX)
		ret = append(ret, bs...)
		solutionCount++
		return solutionCount >= solutionSize
	})
	return string(ret)
}
