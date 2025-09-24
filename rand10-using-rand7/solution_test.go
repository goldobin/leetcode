package rand10_using_rand7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rand7(t *testing.T) {
	n := 7000
	counts := make(map[int]int, n)
	for i := 0; i < n; i++ {
		v := rand7()
		assert.Condition(t, func() bool { return v >= 1 && v <= 7 })
		counts[v]++
	}

	assert.Len(t, counts, 7)
	for _, v := range counts {
		assert.Condition(t, func() bool { return v > 900 && v < 1100 })
	}
}

func Test_rand10(t *testing.T) {
	n := 10000
	counts := make(map[int]int, n)
	for i := 0; i < n; i++ {
		v := rand10()
		assert.Condition(t, func() bool { return v >= 1 && v <= 10 })
		counts[v]++
	}

	assert.Len(t, counts, 10)
	for _, v := range counts {
		assert.Condition(t, func() bool { return v > 900 && v < 1100 })
	}
}
