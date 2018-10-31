package twosum

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6
	expected := []int{1, 2}

	rlt := twoSum(nums, target)

	if !reflect.DeepEqual(rlt, expected) {
		fmt.Println(rlt)
		t.Fail()
	}
}
