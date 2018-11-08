package romanToInt

import (
	"fmt"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	d := map[int]string{
		0:    "",
		3:    "III",
		4:    "IV",
		9:    "IX",
		58:   "LVIII",
		1994: "MCMXCIV",
	}
	for e, s := range d {
		r := romanToInt(s)
		if r != e {
			fmt.Printf("input=%s expected=%d, r=%d\n", s, e, r)
			t.Fail()
		}
	}
}
