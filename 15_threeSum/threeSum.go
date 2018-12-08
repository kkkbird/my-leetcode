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

	var firstZero, lastZero, negetiveEnd, positiveStart = -1, -1, -1, -1

	if nums[0] > 0 || nums[l-1] < 0 { //all positive or all negetive
		return rlt
	}

	for i := 1; i < l; i++ {
		if nums[i-1] < 0 {
			if nums[i] == 0 {
				firstZero = i
				lastZero = i + 1
				negetiveEnd = i
				positiveStart = i + 1
			} else if nums[i] > 0 {
				negetiveEnd = i
				positiveStart = i
			}
		} else if nums[i-1] == 0 {
			if nums[i] == 0 {
				lastZero = i + 1
				positiveStart = i + 1
			} else {
				positiveStart = i
			}
		}
	}

	if lastZero-firstZero >= 3 {
		rlt = append(rlt, []int{0, 0, 0})
	}

	var m map[int]byte
	if firstZero > 0 {
		m = make(map[int]byte)
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
	}

	m = make(map[int]byte)

	for i := 0; i < negetiveEnd; i++ {
		m[-nums[i]] = 0
	}

	for i := positiveStart; i < l-1; i++ {
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

	for i := positiveStart; i < l; i++ {
		m[-nums[i]] = 0
	}

	for i := 0; i < negetiveEnd-1; i++ {
		for j := i + 1; j < negetiveEnd; j++ {
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
