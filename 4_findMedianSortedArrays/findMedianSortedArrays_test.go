package findMedianSortedArrays

import "testing"

func TestIndMedianSortedArrays(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	expected := 2.5

	if findMedianSortedArrays(nums1, nums2) != expected {
		t.Fail()
	}

}
