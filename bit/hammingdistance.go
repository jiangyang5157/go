package hammingdistance

func hammingDistance(a, b uint8) int {
	var ret int = 0
	var i uint
	for i = 0; i < 8; i++ {
		if a&(1<<i) != b&(1<<i) {
			ret++
		}
	}
	return ret
}
