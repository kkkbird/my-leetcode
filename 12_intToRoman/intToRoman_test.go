package intToRoman

import (
	"fmt"
	"testing"
)

func TestIntToRoman(t *testing.T) {
	d := map[int]string{
		0:    "",
		3:    "III",
		4:    "IV",
		9:    "IX",
		58:   "LVIII",
		1994: "MCMXCIV",
	}
	for i, e := range d {
		r := intToRoman(i)
		if r != e {
			fmt.Printf("input=%d expected=%s, r=%s\n", i, e, r)
			t.Fail()
		}
	}
}
