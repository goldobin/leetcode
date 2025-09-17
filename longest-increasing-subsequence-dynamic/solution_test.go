package longest_increasing_subsequence_dynamic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lengthOfLIS(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "case 0.1 empty",
			want: 0,
		},
		{
			name: "case 0.2 one element",
			nums: []int{2},
			want: 1,
		},
		{
			name: "case 0.3 two consecutive index",
			nums: []int{2, 3},
			want: 2,
		},
		{
			name: "case 0.4.1 two non-consecutive index",
			nums: []int{3, 2},
			want: 1,
		},
		{
			name: "case 0.4.2 two equal index",
			nums: []int{2, 2},
			want: 1,
		},
		{
			name: "case 0.5 three consecutive index",
			nums: []int{1, 2, 3},
			want: 3,
		},
		{
			name: "case 0.6.1 three non-consecutive index",
			nums: []int{1, 3, 2},
			want: 2,
		},
		{
			name: "case 0.6.2 three reverse index",
			nums: []int{3, 2, 1},
			want: 1,
		},
		{
			name: "case 0.6.3 three equal index",
			nums: []int{3, 3, 3},
			want: 1,
		},
		{
			name: "case 0.7 four consecutive index",
			nums: []int{1, 2, 3, 4},
			want: 4,
		},
		{
			name: "case 0.8.1 four non-consecutive index",
			nums: []int{1, 3, 2, 4},
			want: 3,
		},
		{
			name: "case 0.8.2 four non-consecutive index",
			nums: []int{1, 3, 2, 4},
			want: 3,
		},
		{
			name: "case 0.8.3 four reverse index",
			nums: []int{4, 3, 2, 1},
			want: 1,
		},
		{
			name: "case 0.8.4 four equal index",
			nums: []int{4, 4, 4, 4},
			want: 1,
		},
		{
			name: "case 1.1 leetcode",
			nums: []int{10, 9, 2, 5, 3, 7, 101, 18},
			want: 4,
		},
		{
			name: "case 1.2 leetcode",
			nums: []int{0, 1, 0, 3, 2, 3},
			want: 4,
		},
		{
			name: "case 1.3 leetcode",
			nums: []int{7, 7, 7, 7, 7, 7, 7},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lengthOfLIS(tt.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}
