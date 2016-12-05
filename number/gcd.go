package number

func GcdBinary(u uint, v uint) uint {
	switch {
	case u == v:
		return u
	case u == 0:
		return v
	case v == 0:
		return u
	}

	uLowestBit := u & 1
	vLowestBit := v & 1
	switch {
	case uLowestBit == 0 && vLowestBit == 0: // u and v are even
		return GcdBinary(u>>1, v>>1) << 1
	case uLowestBit == 0: // u is even, v is odd
		return GcdBinary(u>>1, v)
	case vLowestBit == 0: // u is odd, v is even
		return GcdBinary(u, v>>1)
	case u > v: // u and v are odd, u > v
		return GcdBinary((u-v)>>1, v)
	case u < v: // u and v are odd, u < v
		return GcdBinary(u, (v-u)>>1)
	default: // u and v are odd, u == v
		return u
	}
}

func GcdEuclidean(u uint, v uint) uint {
	switch {
	case u == v:
		return u
	case u == 0:
		return v
	case v == 0:
		return u
	default:
		return GcdEuclidean(v, u%v)
	}
}
