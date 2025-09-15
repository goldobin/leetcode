package summary_ranges

import (
	"fmt"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}

	var (
		firstV               = nums[0]
		result      []string = nil
		appendRange          = func(a int, b int) {
			var s string
			if a == b {
				s = fmt.Sprintf("%d", a)
			} else {
				s = fmt.Sprintf("%d->%d", a, b)
			}

			result = append(result, s)
		}
	)

	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] == 1 {
			continue
		}
		appendRange(firstV, nums[i-1])
		firstV = nums[i]
	}

	appendRange(firstV, nums[len(nums)-1])
	return result
}
