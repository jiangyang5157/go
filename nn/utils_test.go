package nn

import (
	"fmt"
	"testing"
)

func TestUtils_RandGuass(t *testing.T) {
	RandSeed()
	for i := 0; i < 20; i++ {
		fmt.Println(RandGuass(0, 100))
	}
}
