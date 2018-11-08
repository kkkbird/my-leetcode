package longestCommonPrefix

import (
	"math"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	l := math.MaxInt64

	for _, s := range strs {
		l = int(math.Min(float64(l), float64(len(s))))
	}

	rlt := ""
	var c byte
	for i := 0; i < l; i++ {
		c = 0
		for _, s := range strs {
			if c == 0 {
				c = s[i]
			} else {
				if c != s[i] {
					return rlt
				}
			}
		}
		rlt += string(c)
	}

	return rlt
}
