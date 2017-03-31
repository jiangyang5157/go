package hammingdistance

func hammingDistance(a, b uint8) int {
	var ret = 0
	if a&1<<0 != b&1<<0 {
		ret++
	}
	if a&1<<1 != b&1<<1 {
		ret++
	}
	if a&1<<2 != b&1<<2 {
		ret++
	}
	if a&1<<3 != b&1<<3 {
		ret++
	}
	if a&1<<4 != b&1<<4 {
		ret++
	}
	if a&1<<5 != b&1<<5 {
		ret++
	}
	if a&1<<6 != b&1<<6 {
		ret++
	}
	if a&1<<7 != b&1<<7 {
		ret++
	}
	return ret
}
