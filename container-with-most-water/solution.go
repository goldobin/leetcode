package container_with_most_water

func maxArea(height []int) int {
	var (
		l, r = 0, len(height) - 1
		maxA = 0
	)

	for {
		if l >= r {
			return maxA
		}

		h := min(height[l], height[r])
		w := r - l
		area := h * w
		maxA = max(maxA, area)

		if height[l] < height[r] {
			l = l + 1
		} else {
			r = r - 1
		}
	}
}

func maxAreaBruteForce(height []int) int {
	maxA := 0

	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			h := min(height[i], height[j])
			w := j - i
			area := h * w
			maxA = max(maxA, area)
		}
	}

	return maxA
}
