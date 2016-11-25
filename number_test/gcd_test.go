package number_test

import (
	"testing"
	"github.com/jiangyang5157/go/number"
)

func Test_GcdBinary(t *testing.T) {
	if number.GcdBinary(8, 12) != 4 {
		t.Error("GcdBinary(8, 12) is wrong")
	}

	if number.GcdBinary(12, 8) != 4 {
		t.Error("GcdBinary(12, 8) is wrong")
	}

	if number.GcdBinary(12, 16) != 4 {
		t.Error("GcdBinary(12, 16) is wrong")
	}

	if number.GcdBinary(0, 8) != 8 {
		t.Error("GcdBinary(0, 8) is wrong")
	}

	if number.GcdBinary(8, 1) != 1 {
		t.Error("GcdBinary(8, 1) is wrong")
	}

	if number.GcdBinary(8, 4) != 4 {
		t.Error("GcdBinary(8, 4) is wrong")
	}

	if number.GcdBinary(4, 8) != 4 {
		t.Error("GcdBinary(4, 8) is wrong")
	}

	if number.GcdBinary(8, 8) != 8 {
		t.Error("GcdBinary(8, 8) is wrong")
	}
}

func Test_GcdEuclidean(t *testing.T) {
	if number.GcdEuclidean(8, 12) != 4 {
		t.Error("GcdEuclidean(8, 12) is wrong")
	}

	if number.GcdEuclidean(12, 8) != 4 {
		t.Error("GcdEuclidean(12, 8) is wrong")
	}

	if number.GcdEuclidean(12, 16) != 4 {
		t.Error("GcdEuclidean(12, 16) is wrong")
	}

	if number.GcdEuclidean(0, 8) != 8 {
		t.Error("GcdEuclidean(0, 8) is wrong")
	}

	if number.GcdEuclidean(8, 1) != 1 {
		t.Error("GcdEuclidean(8, 1) is wrong")
	}

	if number.GcdEuclidean(8, 4) != 4 {
		t.Error("GcdEuclidean(8, 4) is wrong")
	}

	if number.GcdEuclidean(4, 8) != 4 {
		t.Error("GcdEuclidean(4, 8) is wrong")
	}

	if number.GcdEuclidean(8, 8) != 8 {
		t.Error("GcdEuclidean(8, 8) is wrong")
	}
}