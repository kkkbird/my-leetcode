package findMedianSortedArrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)

	if m > n {
		return findMedianSortedArrays(nums2, nums1)
	}

	iMin := 0
	iMax := m

	halfLen := (m + n + 1) / 2

	var i, j int

	for iMin <= iMax {
		i = (iMin + iMax) / 2
		j = halfLen - i
		if i < iMax && nums2[j-1] > nums1[i] {
			iMin = i + 1
		} else if i > iMin && nums1[i-1] > nums2[j] {
			iMax = i - 1
		} else {
			break
		}
	}

	maxLeft := 0.0
	if i == 0 {
		maxLeft = float64(nums2[j-1])
	} else if j == 0 {
		maxLeft = float64(nums1[i-1])
	} else {
		if nums1[i-1] > nums2[j-1] {
			maxLeft = float64(nums1[i-1])
		} else {
			maxLeft = float64(nums2[j-1])
		}
	}

	if (m+n)%2 == 1 {
		return float64(maxLeft)
	}

	minRight := 0.0
	if i == m {
		minRight = float64(nums2[j])
	} else if j == n {
		minRight = float64(nums1[i])
	} else {
		if nums1[i] > nums2[j] {
			minRight = float64(nums2[j])
		} else {
			minRight = float64(nums1[i])
		}
	}

	return (minRight + maxLeft) / 2.0
}
