package merge_sorted_array

import (
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	if len(nums1) != m+n {
		panic(fmt.Sprintf("the lengt of nums1 should be %d", m+n))
	}

	if n == 0 {
		return
	}

	if m == 0 {
		for j := 0; j < n; j++ {
			nums1[j] = nums2[j]
		}
	}

	i := m - 1
	j := n - 1
	k := len(nums1) - 1
	for k >= 0 && j >= 0 {

		if i < 0 {
			for n := 0; n <= j; n++ {
				nums1[n] = nums2[n]
			}
			break
		}

		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else if nums1[i] < nums2[j] {
			nums1[k] = nums2[j]
			j--
		} else {
			nums1[k] = nums2[j]
			j--
			k--
			nums1[k] = nums1[i]
			i--
		}
		k--
	}
}
