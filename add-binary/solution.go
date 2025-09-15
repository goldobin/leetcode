package add_binary

func addBinary(a string, b string) string {
	if len(a) == 0 && len(b) == 0 {
		return ""
	}

	if len(a) == 0 {
		return b
	}

	if len(b) == 0 {
		return a
	}

	var (
		result  = make([]rune, max(len(a), len(b))+1)
		caryBit = false
	)

	for i := 0; i < len(result); i++ {
		var (
			aIx              = len(a) - i - 1
			bIx              = len(b) - i - 1
			rIx              = len(result) - i - 1
			aBit, bBit, rBit bool
		)

		if aIx >= 0 && aIx < len(a) {
			aBit = a[aIx] == '1'
		}
		if bIx >= 0 && bIx < len(b) {
			bBit = b[bIx] == '1'
		}

		if !aBit && !bBit {
			rBit = caryBit
			caryBit = false
		} else if aBit && bBit {
			rBit = caryBit
			caryBit = true
		} else if caryBit {
			rBit = false
		} else {
			rBit = true
		}

		if rBit {
			result[rIx] = '1'
		} else {
			result[rIx] = '0'
		}
	}

	if len(result) > 1 && result[0] == '0' {
		return string(result[1:])
	}

	return string(result)
}
