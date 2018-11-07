package convert

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	ss := make([]string, numRows)

	for i, c := range s {
		j := i % (2*numRows - 2)
		if j < numRows {
			ss[j] = ss[j] + string(c)
		} else {
			ss[2*numRows-2-j] = ss[2*numRows-2-j] + string(c)
		}
	}

	rlt := ""
	for i := 0; i < numRows; i++ {
		rlt += ss[i]
	}

	return rlt
}
