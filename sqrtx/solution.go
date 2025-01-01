package sqrtx

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}

	if x == 1 {
		return 1
	}

	var (
		l = 1
		r = x
	)

	for {
		if r-l == 1 {
			return l
		}

		mid := l + (r-l)/2
		midSq := mid * mid

		if midSq == x {
			return mid
		}

		if midSq < x {
			l = mid
		} else {
			r = mid
		}
	}
}
