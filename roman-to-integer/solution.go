package roman_to_integer

import (
	"fmt"
)

/**
 * Roman numbers
 *
 *  I 	-> 1
 *  II 	-> 2
 *  III -> 3
 *  IV 	-> 4
 *  V 	-> 5
 *  VI 	-> 6
 *  VII -> 7
 *  VIII-> 8
 *  IX 	-> 9
 *
 *  1 - 10 are starting from I or V
 *
 *  X 		-> 10
 *  XX		-> 20
 *  XXX 	-> 30
 *  XL 	    -> 40
 *  L 		-> 50
 *  LX		-> 60
 *  LXX		-> 70
 *  LXXX	-> 80
 *  XC		-> 90
 *
 *  10 - 100 are starting from X or L
 */

func romanToInt(s string) int {
	r, err := romanToInt1(s)
	if err != nil {
		return 0
	}

	return r
}

type tokenClass struct {
	oneSymbol  byte
	fiveSymbol byte
	nineSymbol byte
	multiplier int
}
type token struct {
	total      int
	tokenClass tokenClass
	closed     bool
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
	tokenClasss100 = tokenClass{
		oneSymbol:  'C',
		fiveSymbol: 'D',
		nineSymbol: 'M',
		multiplier: 100,
	}
	tokenClasss1000 = tokenClass{
		oneSymbol:  'M',
		fiveSymbol: 0,
		nineSymbol: 0,
		multiplier: 1000,
	}
	allTokenClasses = []tokenClass{
		tokenClass1,
		tokenClass10,
		tokenClasss100,
		tokenClasss1000,
	}
)

func newToken(b byte) (token, error) {
	var (
		tc      tokenClass
		tcFound bool
		total   int
	)

	for _, v := range allTokenClasses {
		switch b {
		case v.oneSymbol:
			tcFound = true
			tc = v
			total = 1
			break
		case v.fiveSymbol:
			tcFound = true
			tc = v
			total = 5
			break
		default:
			continue
		}
	}

	if !tcFound {
		return token{}, fmt.Errorf("unsupported roman number %c", b)
	}

	return token{
		total:      total,
		tokenClass: tc,
		closed:     false,
	}, nil
}

func (t *token) update(b byte) (bool, error) {
	if t.closed {
		return false, fmt.Errorf("token is closed")
	}

	if t.total == 1 {
		updated := false
		switch b {
		case t.tokenClass.oneSymbol:
			t.total += 1
			updated = true
		case t.tokenClass.fiveSymbol:
			t.total = 4
			t.closed = true
			updated = true
		case t.tokenClass.nineSymbol:
			t.total = 9
			t.closed = true
			updated = true
		default:
			t.closed = true
		}
		return updated, nil
	}

	if t.total < 3 || (t.total >= 5 && t.total < 8) {
		updated := false
		switch b {
		case t.tokenClass.oneSymbol:
			t.total += 1
			updated = true

		default:
			t.closed = true
		}

		if t.total == 3 || t.total == 8 {
			t.closed = true
		}

		return updated, nil
	}

	panic("unreachable")
}

func romanToInt1(s string) (int, error) {
	if len(s) == 0 {
		return 0, fmt.Errorf("empty string")
	}

	t, err := newToken(s[0])
	if err != nil {
		return 0, err
	}

	total := 0

	for i := 1; i < len(s); i++ {
		updated, err := t.update(s[i])
		if err != nil {
			return 0, err
		}

		if !t.closed {
			continue
		}

		total += t.total * t.tokenClass.multiplier

		if updated {
			i++
			if i >= len(s) {
				break
			}
		}

		nextT, err := newToken(s[i])
		if err != nil {
			return 0, err
		}

		if nextT.tokenClass.multiplier > t.tokenClass.multiplier {
			return 0, fmt.Errorf(
				"number %c cannot appear after %c or %c",
				s[i],
				t.tokenClass.oneSymbol,
				t.tokenClass.fiveSymbol,
			)
		}

		t = nextT
	}

	if !t.closed {
		total += t.total * t.tokenClass.multiplier
	}

	return total, nil
}
