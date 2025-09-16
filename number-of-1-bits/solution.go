package number_of_1_bits

func hammingWeight(n int) int {
	var (
		mask    = 0b1
		counter = 0
	)

	for i := 0; i < 32; i++ {
		if n&mask == 1 {
			counter++
		}
		n = n >> 1
	}

	return counter
}
