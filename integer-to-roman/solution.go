package integer_to_roman

import (
	"fmt"
	"strings"
)

func intToRoman(num int) string {
	if num <= 0 {
		return ""
	}

	if num > 3999 {
		return ""
	}

	var (
		dividers = []int{1000, 100, 10, 1}
		b        strings.Builder
		n        = num
	)

	for _, divider := range dividers {
		m := n / divider
		k := n % divider

		if m == 0 {
			continue
		}

		v := tokenClassFor(divider).roman(m)
		_, err := b.Write(v)
		if err != nil {
			panic(err)
		}

		n = k
	}

	return b.String()
}

type tokenClass struct {
	oneSymbol  byte
	fiveSymbol byte
	nineSymbol byte
	multiplier int
}

func tokenClassFor(m int) tokenClass {
	switch m {
	case 1:
		return tokenClass1
	case 10:
		return tokenClass10
	case 100:
		return tokenClass100
	case 1000:
		return tokenClass1000
	default:
		panic(fmt.Sprintf("unsupported token class: %v", m))
	}
}

func (t tokenClass) roman(n int) []byte {
	var r []byte
	switch n {
	case 1:
		r = []byte{t.oneSymbol}
	case 2:
		r = []byte{t.oneSymbol, t.oneSymbol}
	case 3:
		r = []byte{t.oneSymbol, t.oneSymbol, t.oneSymbol}
	case 4:
		r = []byte{t.oneSymbol, t.fiveSymbol}
	case 5:
		r = []byte{t.fiveSymbol}
	case 6:
		r = []byte{t.fiveSymbol, t.oneSymbol}
	case 7:
		r = []byte{t.fiveSymbol, t.oneSymbol, t.oneSymbol}
	case 8:
		r = []byte{t.fiveSymbol, t.oneSymbol, t.oneSymbol, t.oneSymbol}
	case 9:
		r = []byte{t.oneSymbol, t.nineSymbol}
	default:
		panic(fmt.Sprintf("unsupported number %v", n))
	}
	return r
}

var (
	tokenClass1 = tokenClass{
		oneSymbol:  'I',
		fiveSymbol: 'V',
		nineSymbol: 'X',
		multiplier: 1,
	}
	tokenClass10 = tokenClass{
		oneSymbol:  'X',
		fiveSymbol: 'L',
		nineSymbol: 'C',
		multiplier: 10,
	}
	tokenClass100 = tokenClass{
		oneSymbol:  'C',
		fiveSymbol: 'D',
		nineSymbol: 'M',
		multiplier: 100,
	}
	tokenClass1000 = tokenClass{
		oneSymbol:  'M',
		fiveSymbol: 0,
		nineSymbol: 0,
		multiplier: 1000,
	}
)
