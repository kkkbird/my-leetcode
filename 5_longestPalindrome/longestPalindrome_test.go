package longestPalindrome

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	s := "cbbd"
	expected := "bb"

	if longestPalindrome(s) != expected {
		t.Fail()
	}
}
