package lengthOfLongestSubstring

func lengthOfLongestSubstring(s string) int {
	l := len(s)
	m := make(map[rune]int)

	maxlength := 0

	for i, j := 0, 0; j < l; j++ {
		if idx, ok := m[rune(s[j])]; ok {
			if idx >= i {
				i = idx + 1
			}
		}
		if j-i+1 > maxlength {
			maxlength = j - i + 1
		}
		m[rune(s[j])] = j
	}

	return maxlength
}

func lengthOfLongestSubstring2(s string) int {
	l := len(s)
	m := make([]int, 256)

	maxlength := 0

	for i, j := 0, 0; j < l; j++ {
		if idx := m[s[j]]; idx != 0 {
			if idx > i {
				i = idx
			}
		}
		if j-i+1 > maxlength {
			maxlength = j - i + 1
		}
		m[s[j]] = j + 1
	}

	return maxlength
}
