package merge_sorted_array

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	cases := []struct {
		nums1    []int
		m        int
		nums2    []int
		n        int
		expected []int
	}{
		{
			nums1:    []int{1, 2, 3, 0, 0, 0},
			m:        3,
			nums2:    []int{2, 5, 6},
			n:        3,
			expected: []int{1, 2, 2, 3, 5, 6},
		},
		{
			nums1:    []int{1, 2, 3, 0, 0, 0},
			m:        3,
			nums2:    []int{4, 5, 6},
			n:        3,
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			nums1:    []int{4, 5, 6, 0, 0, 0},
			m:        3,
			nums2:    []int{1, 2, 3},
			n:        3,
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			nums1:    []int{1},
			m:        1,
			nums2:    []int{},
			n:        0,
			expected: []int{1},
		},

		{
			nums1:    []int{0},
			m:        0,
			nums2:    []int{1},
			n:        1,
			expected: []int{1},
		},
		{
			nums1:    []int{1, 2, 0},
			m:        2,
			nums2:    []int{1},
			n:        1,
			expected: []int{1, 1, 2},
		},
		{
			nums1:    []int{1, 0, 0},
			m:        1,
			nums2:    []int{1, 2},
			n:        2,
			expected: []int{1, 1, 2},
		},
		{
			nums1:    []int{1, 0},
			m:        1,
			nums2:    []int{1},
			n:        1,
			expected: []int{1, 1},
		},
	}

	for i, c := range cases {
		caseName := fmt.Sprintf("Case %02d", i+1)
		t.Run(caseName, func(t *testing.T) {
			result := make([]int, len(c.nums1))
			copy(result, c.nums1)
			merge(result, c.m, c.nums2, c.n)

			if !reflect.DeepEqual(result, c.expected) {
				t.Errorf("expected=%d, got=%d", c.expected, result)
			}
		})
	}
}
