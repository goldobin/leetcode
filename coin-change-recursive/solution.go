package coin_change_recursive

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

	r := iterate(coins, amount)
	if r == math.MaxInt {
		return -1
	}
	return r
}

func iterate(coins []int, rest int) int {
	if rest == 0 {
		return 0
	}

	if rest < 0 {
		return math.MaxInt
	}

	if len(coins) == 0 {
		return math.MaxInt
	}

	if len(coins) == 1 {
		h := coins[0]
		if rest%h != 0 {
			return math.MaxInt
		}

		return rest / h
	}

	head := coins[0]
	if head > rest {
		return iterate(coins[1:], rest)
	}

	r1 := iterate(coins, rest-head)
	r2 := iterate(coins[1:], rest)

	if r1 < r2 {
		return r1 + 1
	} else {
		return r2
	}
}
