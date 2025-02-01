package string_to_integer_atoi

import "math"

func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	var (
		bs                     = []byte(s)
		sign             int32 = 1
		acc              int32 = 0
		expectWhitespace       = true
		expectSign             = true
	)

	for i := 0; i < len(bs); i++ {
		c := bs[i]

		if isSpace(c) {
			if !expectWhitespace {
				break
			}
			continue
		}

		if isSign(c) {
			if !expectSign {
				break
			}

			expectSign = false
			expectWhitespace = false
			if c == '-' {
				sign = -1
			}
			continue
		}

		if !isDigit(c) {
			break
		}

		expectWhitespace = false
		expectSign = false

		prevAcc := acc
		acc = acc*10 + ctoi(c)

		if acc/10 != prevAcc {
			if sign == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}
	}

	return int(acc * sign)
}

func isSign(c byte) bool {
	return c == '+' || c == '-'
}

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func isSpace(c byte) bool {
	return c == ' '
}

func ctoi(c byte) int32 {
	return int32(c - '0')
}
