package letterCombinations

import (
	"fmt"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	input := "237"

	r := letterCombinations(input)

	//if !reflect.DeepEqual(r, expected) {
	fmt.Printf("r=%v", r)
	//	t.Fail()
	//}
}
