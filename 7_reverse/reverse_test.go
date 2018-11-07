package reverse

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	x := 340
	expected := 43

	fmt.Println(x, reverse(x), expected)
	if reverse(x) != expected {
		t.Fail()
	}
}
