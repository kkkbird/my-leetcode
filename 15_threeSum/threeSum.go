package threeSum

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	l := len(nums)
	rlt := make([][]int, 0)
	if l < 3 {
		return rlt
	}

	for i := 0; i < l-2; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		for j := i + 1; j < l-1; j++ {
			if j > i+1 && nums[j-1] == nums[j] {
				continue
			}
			for k := j + 1; k < l; k++ {
				if k > j+1 && nums[k-1] == nums[k] {
					continue
				}
				if nums[i]+nums[j]+nums[k] == 0 {
					rlt = append(rlt, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}

	return rlt
}

func threeSum2(nums []int) [][]int {
	sort.Ints(nums)

	rlt := make([][]int, 0)

	l := len(nums)

	if l < 3 {
		return rlt
	}

	var firstZero, lastZero int

	if firstZero = sort.SearchInts(nums, 0); firstZero < len(nums) { //has 0
		lastZero = firstZero + 1

		for lastZero < l && nums[lastZero] == 0 {
			lastZero++
		}
	}

	if lastZero-firstZero >= 3 {
		rlt = append(rlt, []int{0, 0, 0})
	}

	if firstZero == 0 || lastZero == l {
		return rlt
	}

	m := make(map[int]byte)

	for i := 0; i < firstZero; i++ {
		m[-nums[i]] = 0
	}

	for i := lastZero; i < l; i++ {
		if v, ok := m[nums[i]]; ok {
			if v == 0 {
				m[nums[i]] = 1
				rlt = append(rlt, []int{-nums[i], 0, nums[i]})
			}
		}
	}

	m = make(map[int]byte)

	for i := 0; i < firstZero; i++ {
		m[-nums[i]] = 0
	}

	for i := lastZero; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if v, ok := m[nums[i]+nums[j]]; ok {
				if v == 0 {
					m[nums[i]] = 1
					rlt = append(rlt, []int{-(nums[i] + nums[j]), nums[i], nums[j]})
				}
			}
		}
	}

	m = make(map[int]byte)

	for i := lastZero; i < l; i++ {
		m[-nums[i]] = 0
	}

	for i := 0; i < firstZero-1; i++ {
		for j := i + 1; j < firstZero; j++ {
			if v, ok := m[nums[i]+nums[j]]; ok {
				if v == 0 {
					m[nums[i]] = 1
					rlt = append(rlt, []int{nums[i], nums[j], -(nums[i] + nums[j])})
				}
			}
		}
	}

	return rlt
}
