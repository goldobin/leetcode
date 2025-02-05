package plus_one

func plusOne(digits []int) []int {
	ix := len(digits) - 1
	for {
		if ix < 0 {
			return append([]int{1}, digits...)
		}

		if digits[ix] < 9 {
			digits[ix] += 1
			return digits
		}

		digits[ix] = 0
		ix -= 1
	}
}
