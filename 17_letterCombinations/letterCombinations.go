package letterCombinations

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	m := map[rune][2]byte{
		'2': [2]byte{'a', 3},
		'3': [2]byte{'d', 3},
		'4': [2]byte{'g', 3},
		'5': [2]byte{'j', 3},
		'6': [2]byte{'m', 3},
		'7': [2]byte{'p', 4},
		'8': [2]byte{'t', 3},
		'9': [2]byte{'w', 4},
	}

	l := 1

	for i := 0; i < len(digits); i++ {
		l *= int(m[rune(digits[i])][1])
	}

	rlt := make([][]byte, l)

	for j := 0; j < l; j++ {
		rlt[j] = make([]byte, len(digits))

		v := 1
		for i := len(digits) - 1; i >= 0; i-- {
			rlt[j][i] += m[rune(digits[i])][0] + byte(j/v%int(m[rune(digits[i])][1]))
			v *= int(m[rune(digits[i])][1])
		}
	}

	r := make([]string, l)

	for j := 0; j < l; j++ {
		r[j] = string(rlt[j])
	}
	return r
}
