package coin_exchange

import (
	"math"
	"slices"
)

func coinChange(coins []int, amount int) int {
	if amount <= 0 {
		return 0
	}

	m := make(map[int]int)
	cs := slices.Clone(coins)
	slices.SortFunc(cs, reverseInt)

	r := iterate(m, cs, amount)
	if r == math.MaxInt {
		return -1
	}
	//fmt.Printf("memory keys used: %d", len(m))
	return r
}

var reverseInt = func(a, b int) int {
	if a == b {
		return 0
	}

	if a > b {
		return -1
	}
	return 1
}

func iterate(m memory, coins []int, rest int) int {

	if rest == 0 {
		return 0
	}

	if rest < 0 {
		return math.MaxInt
	}

	if len(coins) == 0 {
		return math.MaxInt
	}

	memorize := false
	if len(coins) <= 1 {
		if memorized, ok := m[rest]; ok {
			return memorized
		}

		memorize = true
	}

	result := func() int {
		head := coins[0]
		if head <= rest {
			r1 := iterate(m, coins, rest-head)
			r2 := iterate(m, coins[1:], rest)

			if r1 < r2 {
				return r1 + 1
			} else {
				return r2
			}

		} else {
			return iterate(m, coins[1:], rest)
		}
	}()

	if memorize {
		m[rest] = result
	}

	return result
}

type memory map[int]int

//func (m memory) memorize(rest int, vFn func() int) int {
//	if memorized, ok := m[rest]; ok {
//		return memorized
//	}
//
//	v := vFn()
//	m[rest] = v
//
//	return v
//}
