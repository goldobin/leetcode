package powx_n

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return x
	}

	if n == -1 {
		return 1 / x
	}

	v := myPow(x, n/2)
	return v * v * myPow(x, n%2)
}
