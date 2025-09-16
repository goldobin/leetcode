package reverse_bits

func reverseBits(n int) int {
	var (
		n1         = uint32(n)
		res uint32 = 0b0
	)

	for i := 0; i < 32; i++ {
		bit := n1 & 0b1
		res <<= 1
		res |= bit
		n1 >>= 1
	}

	return int(res)
}
