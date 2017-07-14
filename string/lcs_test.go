package string_test

import (
	"testing"

	"github.com/jiangyang5157/golang-start/string"
)

func Test_Lcs(t *testing.T) {
	if string.Lcs("GGCACCACG", "ACGGCGGATACG") != "GGCAACG" {
		t.Error("LCS of GGCACCACG and ACGGCGGATACG should be GGCAACG")
	}
}
