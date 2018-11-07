package myAtoi

import (
	"fmt"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	s := []string{"words and 987", "-91283472332", "   -42", "4193 with words", " "}
	expected := []int{0, -2147483648, -42, 4193, 0}

	for i := 0; i < len(s); i++ {
		r := myAtoi(s[i])
		if r != expected[i] {
			fmt.Printf("input=%s, expected=%d, rlt=%d\n", s[i], expected[i], r)
			t.Fail()
		}
	}
}
