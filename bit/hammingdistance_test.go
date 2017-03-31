package hammingdistance

import (
	"fmt"
	"testing"
)

func Test_HammingDistance(t *testing.T) {
	a, b := uint8(0), uint8(255)
	result := hammingDistance(a, b)
	fmt.Printf("HammingDistance: %08b and %08b = %d\n", a, b, result)
}
