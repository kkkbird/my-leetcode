package twosum

func twoSum(nums []int, target int) []int {
	l := len(nums)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

func twoSum2(nums []int, target int) []int {
	m := make(map[int]int)
	var i, j, n, t int
	var ok bool
	for i, n = range nums {
		if _, ok = m[n]; !ok {
			m[n] = i
		}

		t = target - n
		if j, ok = m[t]; ok && j != i {
			return []int{j, i}
		}
	}

	return nil
}
