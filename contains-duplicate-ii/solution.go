package contains_duplicate_ii

func containsNearbyDuplicate(nums []int, k int) bool {
	if k >= len(nums) {
		k = len(nums) - 1
	}

	m := make(map[int]int, k)
	for i := 0; i < len(nums); i++ {
		if lastIx, ok := m[nums[i]]; ok {
			if abs(lastIx-i) <= k {
				return true
			}
		}

		m[nums[i]] = i
	}

	return false
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
