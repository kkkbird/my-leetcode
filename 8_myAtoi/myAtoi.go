package myAtoi

import (
	"math"
	"strings"
)

func myWordAtoi(s string, negative bool) int {
	if len(s) == 0 {
		return 0
	}

	rlt := 0
	maxInt := int(math.Pow(2, 31) - 1)

	for _, c := range s {
		if c >= '0' && c <= '9' {
			p := int(c - '0')
			if rlt > maxInt/10 ||
				((rlt == maxInt/10 && p > 7 && negative == false) ||
					(rlt == maxInt/10 && p > 8 && negative == true)) {
				if negative {
					return maxInt + 1
				} else {
					return maxInt
				}
			}
			rlt = rlt*10 + p
		} else {
			break
		}
	}

	return rlt

}

func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	sw := strings.TrimSpace(s)

	if len(sw) == 0 {
		return 0
	}

	// sw = strings.SplitN(sw, " ", 2)[0]

	// if len(sw) == 0 {
	// 	return 0
	// }

	if sw[0] == '-' {
		return -myWordAtoi(sw[1:], true)
	} else if sw[0] == '+' {
		return myWordAtoi(sw[1:], false)
	} else {
		return myWordAtoi(sw, false)
	}
}
