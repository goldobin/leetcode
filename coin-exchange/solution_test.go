package coin_exchange

import (
	"fmt"
	"testing"
)

func Test_coinExchange(t *testing.T) {
	cases := []struct {
		coins    []int
		amount   int
		expected int
	}{
		{
			coins:    []int{1, 2, 5},
			amount:   11,
			expected: 3,
		},
		{
			coins:    []int{1, 2, 5},
			amount:   10,
			expected: 2,
		},
		{
			coins:    []int{1, 2, 5},
			amount:   100,
			expected: 20,
		},
		{
			coins:    []int{2},
			amount:   3,
			expected: -1,
		},
		{
			coins:    []int{2},
			amount:   1,
			expected: -1,
		},
		{
			coins:    []int{1},
			amount:   0,
			expected: 0,
		},
		{
			coins:    []int{3, 7, 405, 436},
			amount:   8839,
			expected: 25,
		},
		{
			coins:    []int{411, 412, 413, 414, 415, 416, 417, 418, 419, 420, 421, 422},
			amount:   9864,
			expected: 24,
		},
		{
			coins:    []int{227, 99, 328, 299, 42, 322},
			amount:   9847,
			expected: 31,
		},
		{
			coins:    []int{346, 29, 395, 188, 155, 109},
			amount:   9401,
			expected: 26,
		},
		{
			coins:    []int{431, 62, 88, 428},
			amount:   9084,
			expected: 26,
		},
	}

	for i, c := range cases {
		caseName := fmt.Sprintf("Case %02d", i+1)
		t.Run(caseName, func(t *testing.T) {

			result := coinChange(c.coins, c.amount)

			if result != c.expected {
				t.Errorf("expected=%d, got=%d", c.expected, result)
			}
		})
	}
}
