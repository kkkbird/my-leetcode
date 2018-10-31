package addtwonumbers

import (
	"fmt"
	"reflect"
	"testing"
)

func slice2Node(h []int) *ListNode {
	if len(h) == 0 {
		return nil
	}

	rlt := &ListNode{0, nil}
	var p = rlt

	for _, v := range h {
		p.Next = &ListNode{v, nil}
		p = p.Next
	}

	return rlt.Next
}

func node2Slice(l *ListNode) []int {
	r := []int{}
	for p := l; p != nil; p = p.Next {
		r = append(r, p.Val)
	}

	return r
}

func TestAddTwoNumbers(t *testing.T) {
	// l1 := []int{2, 4, 3}
	// l2 := []int{5, 6, 4}
	// expected := []int{7, 0, 8}
	l1 := []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	l2 := []int{5, 6, 4}
	expected := []int{6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}

	rlt := node2Slice(addTwoNumbers(slice2Node(l1), slice2Node(l2)))

	if !reflect.DeepEqual(rlt, expected) {
		fmt.Println(rlt)
		t.Fail()
	}
}
