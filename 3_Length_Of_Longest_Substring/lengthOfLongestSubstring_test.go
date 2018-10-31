package lengthOfLongestSubstring

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	s := "abcabcbb"
	expect := 3
	if lengthOfLongestSubstring2(s) != expect {
		t.Fail()
	}
}
