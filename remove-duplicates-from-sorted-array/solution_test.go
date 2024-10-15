package remove_duplicates_from_sorted_array

import (
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	cases := []struct {
		nums         []int
		expectedLen  int
		expectedNums []int
	}{
		{
			nums:         []int{1},
			expectedLen:  1,
			expectedNums: []int{1},
		},
		{
			nums:         []int{1, 2},
			expectedLen:  2,
			expectedNums: []int{1, 2},
		},
		{
			nums:         []int{1},
			expectedLen:  1,
			expectedNums: []int{1},
		},
		{
			nums:         []int{1, 1, 2},
			expectedLen:  2,
			expectedNums: []int{1, 2},
		},
		{
			nums:         []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expectedLen:  5,
			expectedNums: []int{0, 1, 2, 3, 4},
		},
	}

	for i, c := range cases {
		l := removeDuplicates(c.nums)

		if l != c.expectedLen {
			t.Errorf("case %02d: expected len=%d, got len=%d", i, c.expectedLen, l)
		}

		for k := 0; k < l; k++ {
			a := c.expectedNums[k]
			b := c.nums[k]
			if a != b {
				t.Errorf("case %02d: pos=%d, expected char=%v, got char=%v", i, k, a, b)
				break
			}
		}

		if l != c.expectedLen {
			t.Errorf("case %02d: expected len=%d, got len=%d", i, c.expectedNums, l)
		}
	}
}
