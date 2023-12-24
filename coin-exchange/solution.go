package coin_exchange

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

func coinChange(coins []int, amount int) int {
	a := newAlg(coins)
	return a.run(amount)
}

type alg struct {
	coins []int
}

func newAlg(coins []int) alg {
	return alg{
		coins: slices.Clone(coins),
	}
}

func (a alg) run(amount int) int {
	if amount <= 0 {
		return 0
	}

	result := a.iterate(
		make(map[string]int),
		newCoinCounts(a.coins),
		amount,
	)

	if result == math.MaxInt {
		return -1
	}

	return result
}

func makeKey(cs []int) string {
	b := strings.Builder{}

	for _, v := range cs {
		b.WriteString(strconv.Itoa(v))
		b.WriteString(",")
	}

	return b.String()
}
func (a alg) iterate(m memory, cc coinCounts, rest int) int {
	if rest == 0 {
		return cc.total
	}

	if rest < 0 {
		return math.MaxInt
	}

	bestCount := math.MaxInt
	for i, v := range a.coins {
		newCC := cc.increment(i, 1)
		memoryKey := makeKey(newCC.totalPerCoin)

		newBestCount := m.memorize(memoryKey, func() int {
			return a.iterate(m, newCC, rest-v)
		})

		if newBestCount >= bestCount {
			continue
		}

		bestCount = newBestCount
	}

	return bestCount
}

type coinCounts struct {
	totalPerCoin []int
	total        int
}

func newCoinCounts(coins []int) coinCounts {
	return coinCounts{
		totalPerCoin: make([]int, len(coins)),
		total:        0,
	}

}

func (cc coinCounts) increment(i, v int) coinCounts {
	newPerCoin := slices.Clone(cc.totalPerCoin)
	newPerCoin[i] += v
	newTotal := cc.total + v

	return coinCounts{
		totalPerCoin: newPerCoin,
		total:        newTotal,
	}
}

type memory map[string]int

func (m memory) memorize(key string, vFn func() int) int {
	if memorized, ok := m[key]; ok {
		return memorized
	}

	v := vFn()
	m[key] = v

	return v
}
