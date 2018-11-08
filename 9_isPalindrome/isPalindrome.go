package isPalindrome

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x == 0 {
		return true
	}

	rlt := 0

	for j := x; j != 0; j = j / 10 {
		p := j % 10
		rlt = rlt*10 + p
	}

	return rlt == x
}
