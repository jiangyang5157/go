package number

import (
	"testing"
)

func Test_GcdBinary(t *testing.T) {
	if GcdBinary(8, 12) != 4 {
		t.Error("GcdBinary(8, 12) is wrong")
	}

	if GcdBinary(12, 8) != 4 {
		t.Error("GcdBinary(12, 8) is wrong")
	}

	if GcdBinary(12, 16) != 4 {
		t.Error("GcdBinary(12, 16) is wrong")
	}

	if GcdBinary(0, 8) != 8 {
		t.Error("GcdBinary(0, 8) is wrong")
	}

	if GcdBinary(8, 1) != 1 {
		t.Error("GcdBinary(8, 1) is wrong")
	}

	if GcdBinary(8, 4) != 4 {
		t.Error("GcdBinary(8, 4) is wrong")
	}

	if GcdBinary(4, 8) != 4 {
		t.Error("GcdBinary(4, 8) is wrong")
	}

	if GcdBinary(8, 8) != 8 {
		t.Error("GcdBinary(8, 8) is wrong")
	}
}

func Test_GcdEuclidean(t *testing.T) {
	if GcdEuclidean(8, 12) != 4 {
		t.Error("GcdEuclidean(8, 12) is wrong")
	}

	if GcdEuclidean(12, 8) != 4 {
		t.Error("GcdEuclidean(12, 8) is wrong")
	}

	if GcdEuclidean(12, 16) != 4 {
		t.Error("GcdEuclidean(12, 16) is wrong")
	}

	if GcdEuclidean(0, 8) != 8 {
		t.Error("GcdEuclidean(0, 8) is wrong")
	}

	if GcdEuclidean(8, 1) != 1 {
		t.Error("GcdEuclidean(8, 1) is wrong")
	}

	if GcdEuclidean(8, 4) != 4 {
		t.Error("GcdEuclidean(8, 4) is wrong")
	}

	if GcdEuclidean(4, 8) != 4 {
		t.Error("GcdEuclidean(4, 8) is wrong")
	}

	if GcdEuclidean(8, 8) != 8 {
		t.Error("GcdEuclidean(8, 8) is wrong")
	}
}
