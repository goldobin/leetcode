package contains_duplicate_ii

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func containsNearbyDuplicate(nums []int, k int) bool {
	for i, v := range nums {
		for j, w := range nums {
			if v == w && i != j && abs(i-j) <= k {
				return true
			}
		}
	}

	return false
}
