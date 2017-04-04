package hammingdistance

import "math/big"

func HammingDistance(a, b int64) int {
	return CountBit(big.NewInt(a ^ b))
}

func CountBit(n *big.Int) int {
	var count int
	for _, x := range n.Bits() {
		for x != 0 {
			x &= x - 1
			count++
		}
	}
	return count
}
