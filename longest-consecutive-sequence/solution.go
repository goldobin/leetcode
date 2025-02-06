package longest_consecutive_sequence

func longestConsecutive(nums []int) int {
	m := make(map[int]struct{}, len(nums))
	longestStreak := 0

	for _, n := range nums {
		m[n] = struct{}{}
	}

	for n := range m {
		if _, ok := m[n-1]; ok {
			continue
		}
		streak := 0
		for {
			if _, ok := m[n]; !ok {
				if longestStreak < streak {
					longestStreak = streak
				}
				break
			}

			n += 1
			streak += 1
		}
	}

	return longestStreak
}
