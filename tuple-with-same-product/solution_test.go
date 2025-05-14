package tuple_with_same_product

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_tupleSameProduct(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "case 0.1",
			nums: []int{2, 3, 4},
			want: 0,
		},
		{
			name: "case 0.2",
			nums: []int{2, 3},
			want: 0,
		},
		{
			name: "case 0.3",
			nums: []int{2},
			want: 0,
		},
		{
			name: "case 0.4",
			nums: []int{},
			want: 0,
		},
		{
			name: "case 1.1 leetcode",
			nums: []int{2, 3, 4, 6},
			want: 8,
		},
		{
			name: "case 1.2 leetcode",
			nums: []int{1, 2, 4, 5, 10},
			want: 16,
		},
		{
			name: "case 1.3",
			nums: []int{1, 2, 4, 5, 10, 20},
			/**
			 * {1, 20, 2, 10} * 8
			 * {1, 10, 2, 5} * 8
			 * {2, 20, 4, 10} * 8f
			 * {2, 10, 4, 5} * 8
			 */
			want: 40,
		},
		{
			name: "case 1.4 leetcode",
			nums: []int{2, 3, 4, 6, 8, 12},
			/**
			 * {2, 12, 3, 8} * 8
			 * {2, 12, 4, 6} * 8
			 * {3, 8, 4, 6} * 8
			 * {4, 12, 6, 8} * 8
			 * {2, 6, 3, 4} * 8
			 */
			want: 40,
		},
		{
			name: "case 1.5 leetcode",
			nums: []int{20, 10, 6, 24, 15, 5, 4, 30},
			want: 48,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tupleSameProduct(tt.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}
