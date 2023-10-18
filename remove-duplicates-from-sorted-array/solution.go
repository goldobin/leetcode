package remove_duplicates_from_sorted_array

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	k := 1
	for i := 1; i < len(nums); i++ {
		a := nums[i-1]
		b := nums[i]

		if a != b {
			nums[k] = b
			k = k + 1
		}
	}

	return k
}
