package rotate_array

import (
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

func Test_rotate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{
			name: "case 0.1",
			nums: []int{},
			k:    0,
			want: []int{},
		},
		{
			name: "case 0.2",
			nums: []int{1},
			k:    0,
			want: []int{1},
		},
		{
			name: "case 0.3",
			nums: []int{1},
			k:    1,
			want: []int{1},
		},
		{
			name: "case 0.4",
			nums: []int{1, 2},
			k:    1,
			want: []int{2, 1},
		},
		{
			name: "case 1.1 leetcode",
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			k:    3,
			want: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name: "case 1.2",
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8},
			k:    3,
			want: []int{6, 7, 8, 1, 2, 3, 4, 5},
		},
		{
			name: "case 1.3",
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8},
			k:    5,
			want: []int{4, 5, 6, 7, 8, 1, 2, 3},
		},
		{
			name: "case 1.4",
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8},
			k:    8,
			want: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name: "case 2.1 leetcode",
			nums: []int{-1, -100, 3, 99},
			k:    2,
			want: []int{3, 99, -1, -100},
		},
		{
			name: "case 2.2",
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8},
			k:    4,
			want: []int{5, 6, 7, 8, 1, 2, 3, 4},
		},
		{
			name: "case 2.3",
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8},
			k:    2,
			want: []int{7, 8, 1, 2, 3, 4, 5, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := slices.Clone(tt.nums)
			rotate(got, tt.k)
			assert.Equal(t, tt.want, got)
		})
	}
}
