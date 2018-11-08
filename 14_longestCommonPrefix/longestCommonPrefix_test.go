package longestCommonPrefix

import (
	"fmt"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	d := map[string][]string{
		"":   []string{},
		"fl": []string{"flower", "flow", "flight"},
		//"":   []string{"dog", "racecar", "car"},
	}
	for e, strs := range d {
		r := longestCommonPrefix(strs)
		if r != e {
			fmt.Printf("input=%v, expected=%s, r=%s\n", strs, e, r)
			t.Fail()
		}
	}
}
