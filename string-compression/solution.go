package string_compression

import "strconv"

func compress(chars []byte) int {
	if len(chars) == 0 {
		return 0
	}

	var (
		startCh      = chars[0]
		startChIndex = 0
		compLen      = 0
	)

	for i := 1; i <= len(chars); i++ {
		var ch byte
		if i < len(chars) {
			ch = chars[i]
		}

		if ch == startCh {
			continue
		}

		repCount := i - startChIndex
		chars[compLen] = startCh
		compLen += 1
		startCh = ch
		startChIndex = i

		if repCount == 1 {
			continue
		}

		repCountCompLen := 1 + log10i(repCount)
		strconv.AppendInt(chars[:compLen], int64(repCount), 10)
		compLen += repCountCompLen
	}

	return compLen
}

func log10i(i int) int {
	if i < 0 {
		panic("log of negative number is undefined")
	}
	n := 0
	for i >= 10 {
		i = i / 10
		n++
	}
	return n
}
