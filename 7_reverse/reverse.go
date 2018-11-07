package reverse

import (
	"math"
)

func reverse(x int) int {
	if x == 0 {
		return 0
	}
	if x < 0 {
		return -reverse(-x)
	}

	maxInt := int(math.Pow(2, 31) - 1)

	rlt := 0

	for j := x; j != 0; j = j / 10 {
		p := j % 10
		if rlt > maxInt/10 || (rlt == maxInt/10 && p > 7) {
			return 0
		}
		rlt = rlt*10 + p
	}
	return rlt
}
