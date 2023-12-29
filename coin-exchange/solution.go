package coin_exchange

import (
	"math"
	"slices"
)

func coinChange(coins []int, amount int) int {
	if amount <= 0 {
		return 0
	}

	cs := slices.Clone(coins)
	slices.SortFunc(cs, reverseInt)

	r := iterate(cs, 0, amount)
	if r == math.MaxInt {
		return -1
	}
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

func iterate(coins []int, count int, rest int) int {

	if rest == 0 {
		return count
	}

	if rest < 0 {
		return math.MaxInt
	}

	if len(coins) == 0 {
		return math.MaxInt
	}

	head := coins[0]

	if head <= rest {
		r1 := iterate(coins, count+1, rest-head)
		r2 := iterate(coins[1:], count, rest)

		if r1 < r2 {
			return r1
		} else {
			return r2
		}
	} else {
		return iterate(coins[1:], count, rest)
	}
}

//bestCount := math.MaxInt
//for i, v := range a.coins {
//	newCC := cc.increment(i, 1)
//	memoryKey := makeKey(newCC.totalPerCoin)
//
//	newBestCount := m.memorize(memoryKey, func() int {
//		return a.iterate(m, newCC, rest-v)
//	})
//
//	if newBestCount >= bestCount {
//		continue
//	}
//
//	bestCount = newBestCount
//}
//
//return bestCount

//type alg struct {
//	coins []int
//}

//func newAlg(coins []int) alg {
//	cs := slices.Clone(coins)
//	slices.SortFunc(cs, reverseInt)
//
//	return alg{
//		coins: cs,
//	}
//}

//func (a alg) run(amount int) int {
//	if amount <= 0 {
//		return 0
//	}
//
//	result := a.iterate(
//		make(map[string]int),
//		newCoinCounts(a.coins),
//		amount,
//	)
//
//	if result == math.MaxInt {
//		return -1
//	}
//
//	return result
//}

//func makeKey(cs []int) string {
//	b := strings.Builder{}
//
//	for _, v := range cs {
//		b.WriteString(strconv.Itoa(v))
//		b.WriteString(",")
//	}
//
//	return b.String()
//}

//type coinCounts struct {
//	totalPerCoin []int
//	total        int
//}
//
//func newCoinCounts(coins []int) coinCounts {
//	return coinCounts{
//		totalPerCoin: make([]int, len(coins)),
//		total:        0,
//	}
//
//}
//
//func (cc coinCounts) increment(i, v int) coinCounts {
//	newPerCoin := slices.Clone(cc.totalPerCoin)
//	newPerCoin[i] += v
//	newTotal := cc.total + v
//
//	return coinCounts{
//		totalPerCoin: newPerCoin,
//		total:        newTotal,
//	}
//}
//
//type memory map[string]int
//
//func (m memory) memorize(key string, vFn func() int) int {
//	if memorized, ok := m[key]; ok {
//		return memorized
//	}
//
//	v := vFn()
//	m[key] = v
//
//	return v
//}
