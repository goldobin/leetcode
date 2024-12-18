package factorial_trailing_zeroes

func trailingZeroes(n int) int {
	a := n / 5
	fs := 0

	for i := 1; i <= a; i++ {
		fs += fives(i * 5)
	}

	return fs
}

func fives(n int) int {
	if n < 5 {
		return 0
	}

	r := n % 5
	if r != 0 {
		return 0
	}

	return 1 + fives(n/5)
}
