package isPalindrome

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	input := []int{121, -121, 10}
	expected := []bool{true, false, false}

	for i := 0; i < len(input); i++ {
		r := isPalindrome(input[i])
		if r != expected[i] {
			fmt.Printf("input=%d, expected=%v, rlt=%v\n", input[i], expected[i], r)
			t.Fail()
		}
	}
}
