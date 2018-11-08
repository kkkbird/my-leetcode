package romanToInt

var romeMap = map[int]string{
	1:    "I",
	5:    "V",
	10:   "X",
	50:   "L",
	100:  "C",
	500:  "D",
	1000: "M",
	4:    "IV",
	9:    "IX",
	40:   "XL",
	90:   "XC",
	400:  "CD",
	900:  "CM",
}

var romeValue = []int{
	1000,
	900,
	500,
	400,
	100,
	90,
	50,
	40,
	10,
	9,
	5,
	4,
	1,
}

func romanToInt(s string) int {
	if len(s) == 0 {
		return 0
	}
	rome2IntMap := make(map[string]int)
	for i, s := range romeMap {
		rome2IntMap[s] = i
	}
	rlt := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) {
			if v, ok := rome2IntMap[s[i:i+2]]; ok {
				rlt += v
				i++
				continue
			}
		}
		rlt += rome2IntMap[s[i:i+1]]
	}

	return rlt
}
