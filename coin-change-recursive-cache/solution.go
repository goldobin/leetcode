package coin_change_recursive_cache

import (
	"math"
	"slices"
)

func coinChange(coins []int, amount int) int {
	if amount <= 0 {
		return 0
	}

	slices.Sort(coins)
	slices.Reverse(coins)

	m := make(map[memKey]int)

	r := iterate(m, coins, amount)
	if r == math.MaxInt {
		return -1
	}
	return r
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

	head := coins[0]
	if len(coins) == 1 {
		if rest%head != 0 {
			return math.MaxInt
		}

		return rest / head
	}

	k := memKey{head, rest}
	if v, ok := m[k]; ok {
		return v
	}

	var v int
	if head > rest {
		v = iterate(m, coins[1:], rest)
	} else {
		r1 := iterate(m, coins, rest-head)
		r2 := iterate(m, coins[1:], rest)

		if r1 < r2 {
			v = r1 + 1
		} else {
			v = r2
		}
	}

	m[k] = v
	return v
}

type (
	memKey struct {
		maxCoin int
		rest    int
	}
	memory map[memKey]int
)
