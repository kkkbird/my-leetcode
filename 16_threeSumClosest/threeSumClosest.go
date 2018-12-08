package threeSumClosest

import (
	"math"
	"sort"
)

func absInt(a int) int {
	return int(math.Abs(float64(a)))
}

func threeSumClosest(nums []int, target int) int {
	rlt := 0
	if len(nums) < 3 {
		return rlt
	}
	sort.Ints(nums)

	rlt = nums[0] + nums[1] + nums[2]

	for i := 0; i < len(nums); i++ {
		j, k := 0, len(nums)-1

		for j < k {
			if j == i || k == i {
				break
			}
			if absInt(rlt-target) > absInt(nums[j]+nums[k]+nums[i]-target) {
				rlt = nums[j] + nums[k] + nums[i]
			}
			if rlt == target {
				return rlt
			} else if nums[j]+nums[k] < target-nums[i] {
				j++
			} else {
				k--
			}
		}
	}

	return rlt
}
