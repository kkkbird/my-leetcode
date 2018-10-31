package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var rlt *ListNode = &ListNode{
		-1,
		nil,
	}

	var current = rlt

	p1 := l1
	p2 := l2
	var a, v, v1, v2 int

	a = 0
	for {
		v = a

		if p1 != nil {
			v1 = p1.Val
		} else {
			v1 = 0
		}

		if p2 != nil {
			v2 = p2.Val
		} else {
			v2 = 0
		}

		v += v1 + v2

		if v >= 10 {
			v -= 10
			a = 1
		} else {
			a = 0
		}

		node := ListNode{
			Val:  v,
			Next: nil,
		}

		if p1 == nil && p2 == nil {
			//done
			if v > 0 {
				current.Next = &node
				current = current.Next
			}
			break
		}

		current.Next = &node
		current = current.Next

		if p1 != nil {
			p1 = p1.Next
		}

		if p2 != nil {
			p2 = p2.Next
		}
	}

	return rlt.Next
}

func int2Node(i int) *ListNode {
	var rlt *ListNode = &ListNode{
		0,
		nil,
	}

	var current = rlt

	for l := i; l > 0; l = l / 10 {
		current.Next = &ListNode{l % 10, nil}
		current = current.Next
	}

	if current == rlt {
		return rlt
	}

	return rlt.Next
}

func node2Int(l *ListNode) (rlt int) {
	seed := 1
	for current := l; current != nil; current = current.Next {
		rlt += current.Val * seed
		seed *= 10
	}
	return
}

//无法通过，有溢出错误
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	return int2Node(node2Int(l1) + node2Int(l2))
}
