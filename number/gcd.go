package number

func GcdBinary(u uint, v uint) uint {
	if u == v {
		return u
	}
	if v == 0 {
		return u
	}
	if u == 0 {
		return v
	}

	uLowestBit := u & 1;
	vLowestBit := v & 1;
	if uLowestBit == 0 && vLowestBit == 0 {
		// u and v are even
		return GcdBinary(u >> 1, v >> 1) << 1
	} else if uLowestBit == 0 {
		// u is even, v is odd
		return GcdBinary(u >> 1, v)
	} else if vLowestBit == 0 {
		// u is odd, v is even
		return GcdBinary(u, v >> 1)
	} else if (u >= v) {
		// u and v are odd, u >= v
		return GcdBinary((u - v) >> 1, v)
	} else {
		// u and v are odd, u < v
		return GcdBinary(u, (v - u) >> 1)
	}
}

func GcdEuclidean(u uint, v uint) uint {
	if v > 0 {
		return GcdEuclidean(v, u % v)
	} else {
		return u
	}
}