package hammingdistance

import (
	"fmt"
	"math/big"
	"testing"
)

func Test_HammingDistance(t *testing.T) {
	var a, b int64 = 1, 7
	result := HammingDistance(a, b)
	fmt.Printf("%064b\n%064b\nHammingDistance = %d\n", a, b, result)
}

func Test_CountBit(t *testing.T) {
	testCases := []struct {
		in  *big.Int
		out int
	}{
		{big.NewInt(0), 0},
		{big.NewInt(1), 1},
		{big.NewInt(2), 1},
		{big.NewInt(3), 2},
		{big.NewInt(4), 1},
		{big.NewInt(123456789), 16},
	}
	for _, testCase := range testCases {
		if CountBit(testCase.in) != testCase.out {
			t.Errorf("CountBit(%v): expect %v, actual %v\n", testCase.in, testCase.out, CountBit(testCase.in))
		}
	}
}
