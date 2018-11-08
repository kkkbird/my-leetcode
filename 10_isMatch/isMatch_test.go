package isMatch

import (
	"fmt"
	"testing"
)

func TestIsMatch(t *testing.T) {
	input1 := []string{"aasdfasdfasdfasdfas", "a", "a", "aaa", "", "aaa", "aaa", "aab", "aabc", "aa", "aa", "ab", "aab", "mississippi", "", "", "a"}
	input2 := []string{"aasdf.*asdf.*asdf.*asdf.*s", ".*.", ".*..a*", "aaaa", ".", "ab*a*c*a", "a*a", "a*.", "a*.", "a", "a*", ".*", "c*a*b", "mis*is*p*.", "a*", "", ""}
	expected := []bool{true, true, false, false, false, true, true, true, false, false, true, true, true, false, true, true, false}

	for i := 0; i < len(input1); i++ {
		r := isMatch(input1[i], input2[i])
		if r != expected[i] {
			fmt.Printf("input1=%s, input2=%s, expected=%v, rlt=%v\n", input1[i], input2[i], expected[i], r)
			t.Fail()
		}
	}
}
