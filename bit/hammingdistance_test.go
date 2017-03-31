package hammingdistance

import (
	"fmt"
	"testing"
)

func Test_HamDist(t *testing.T) {
	var a, b uint8 = 0, 255
	result := hammingDistance(a, b)
	fmt.Printf("HammingDistance: %08b and %08b = %d\n", byte(a), byte(b), result)
}
