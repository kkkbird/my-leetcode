package longestPalindrome

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	if len(s) == 1 {
		return s
	}

	s2 := ""

	for i := range s {
		s2 = s[i:i+1] + s2
	}

	l := len(s)
	maxLength := 0
	var maxIdx, maxCur int

	for idx := 0; idx < l && idx < l-maxLength; idx++ {
		for cur := l; cur > idx+maxLength; cur-- {
			if s[idx:cur] == s2[l-cur:l-idx] && maxLength < cur-idx {
				maxLength = cur - idx
				maxIdx = idx
				maxCur = cur
			}
		}
	}

	return s[maxIdx:maxCur]
}
