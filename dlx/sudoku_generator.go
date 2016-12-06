package dlx

import (
	"fmt"
)

// Generate unique solution puzzle
func GeneratePuzzle(squareLength int, minTotalGivens int, minSubGivens int, maxSubGivens int) string {
	if squareLength < 1 {
		return "Invalid puzzle"
	}

	p := newPuzzle(squareLength)
	var rtp []int = p.randomTerminalPattern()



	//todo
	fmt.Println(rtp)
	return ""
}

func (p *puzzle) randomTerminalPattern() []int {
	// fill diagonal squares by disorder digits
	var ret []int = make([]int, p.cellSize)
	var digits []int = make([]int, p.cellSize)
	var tmp []int = make([]int, p.edgeLength)
	for i := range tmp {
		tmp[i] = i + 1
	}
	// some times the random number for squares cause zero solution, particularly 2x2 puzzle
	for ok := false; ok != true; {
		for i := 0; i < p.edgeLength; i += p.squareLength + 1 {
			var d []int = disorderArray(tmp)
			for j := 0; j < p.edgeLength; j++ {
				r := j / p.squareLength + (i / p.squareLength) * p.squareLength
				c := j % p.squareLength + (i / p.squareLength) * p.squareLength
				digits[p.getCellIndex(r, c)] = d[j];
			}
		}

		// solve
		p.init(digits)
		p.search(func(o []*x) bool {
			for _, o := range p.o {
				x0 := o.x0
				x0ci := x0.c.i
				x0rci := x0.r.c.i
				ret[x0ci - 1] = (x0rci - 1) % p.edgeLength + 1
			}
			ok = true
			return true
		})
	}
	return ret
}
