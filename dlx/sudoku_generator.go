package dlx

func minTotalGivens(edgeLength int) int {
	return edgeLength * 2 - 1 // unconfirmed
}

// Generate unique solution puzzle
func GeneratePuzzle(squareLength int, minSubGivens int, maxSubGivens int) string {
	if squareLength < 1 {
		return "Invalid puzzle"
	}

	p := newPuzzle(squareLength)
	var rtp []int = p.randomTerminalPattern()

	//mtg := minTotalGivens(p.edgeLength)
	//squareLength := p.squareLength
	//edgeLength := p.edgeLength
	//cellSize := p.cellSize
	//var remainRowGivens []int = make([]int, edgeLength)
	//var remainColumnGivens []int = make([]int, edgeLength)



	return digits2raw(rtp)
}

func (p *puzzle) randomTerminalPattern() []int {
	squareLength := p.squareLength
	edgeLength := p.edgeLength
	cellSize := p.cellSize
	var ret []int = make([]int, cellSize)
	var digits []int = make([]int, cellSize)
	var tmp []int = make([]int, edgeLength)
	for i := range tmp {
		tmp[i] = i + 1
	}

	// for-loop: some times the random number for squares cause zero solution, particularly 2x2 puzzle
	for ok := false; ok != true; {
		// fill diagonal squares by disorder digits
		for i := 0; i < edgeLength; i += squareLength + 1 {
			var d []int = disorderArray(tmp)
			for j := 0; j < edgeLength; j++ {
				r := j / squareLength + (i / squareLength) * squareLength
				c := j % squareLength + (i / squareLength) * squareLength
				digits[p.cellIndex(r, c)] = d[j];
			}
		}

		// search for a solution
		p.init(digits)
		p.search(func(o []*x) bool {
			for _, x := range o {
				x0 := x.x0
				x0ci := x0.c.i    // x0ci = [offset1 + 1, offset2]
				x0rci := x0.r.c.i // x0rci = [offset2 + 1, offset3]
				// append by raw, index = [0, cellSize - 1]
				ret[x0ci - 1] = (x0rci - 1) % edgeLength + 1
			}
			ok = true
			return true
		})
	}
	return ret
}
