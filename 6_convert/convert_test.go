package convert

import (
	"testing"
)

func TestConvert(t *testing.T) {
	//s := "PAYPALISHIRING"
	// numRows := 3
	// expected := "PAHNAPLSIIGYIR"

	// numRows := 4
	// expected := "PINALSIGYAHRPI"

	s := "A"
	numRows := 1
	expected := "A"

	if convert(s, numRows) != expected {
		t.Fail()
	}
}
