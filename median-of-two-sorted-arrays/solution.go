package median_of_two_sorted_arrays

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) == 0 {
		panic("nums1 and nums2 are empty")
	}

	if len(nums1) == 0 {
		return median(nums2)
	}

	if len(nums2) == 0 {
		return median(nums1)
	}

	var (
		nums1Len = len(nums1)
		nums2Len = len(nums2)
	)

	if nums1Len > nums2Len {
		nums1, nums2 = nums2, nums1
		nums1Len, nums2Len = nums2Len, nums1Len
	}

	var (
		totalLen = nums1Len + nums2Len
		isOdd    = totalLen%2 != 0
		l        = 0
		r        = nums1Len
	)

	for l <= r {
		var (
			pivot1        = (l + r) / 2
			pivot2        = (nums1Len+nums2Len+1)/2 - pivot1
			nums1LeftMax  = math.MinInt
			nums1RightMin = math.MaxInt
			nums2LeftMax  = math.MinInt
			nums2RightMin = math.MaxInt
		)

		if pivot1 > 0 {
			nums1LeftMax = nums1[pivot1-1]
		}
		if pivot1 < nums1Len {
			nums1RightMin = nums1[pivot1]
		}
		if pivot2 > 0 {
			nums2LeftMax = nums2[pivot2-1]
		}
		if pivot2 < nums2Len {
			nums2RightMin = nums2[pivot2]
		}

		if nums1LeftMax <= nums2RightMin && nums2LeftMax <= nums1RightMin {
			if isOdd {
				return float64(max(nums1LeftMax, nums2LeftMax))
			}

			maxL := max(nums1LeftMax, nums2LeftMax)
			minR := min(nums1RightMin, nums2RightMin)
			return float64(maxL+minR) / 2.0
		}

		if nums1LeftMax > nums2RightMin {
			r = pivot1 - 1
		} else {
			l = pivot1 + 1
		}
	}

	return 0.0
}

func median(nums []int) float64 {
	numsLen := len(nums)
	if numsLen == 0 {
		panic("nums is empty")
	}

	if numsLen == 1 {
		return float64(nums[0])
	}

	m := numsLen / 2
	if numsLen%2 == 0 {
		l := m - 1
		return float64(nums[l]+nums[m]) / 2.0
	}

	return float64(nums[m])
}
