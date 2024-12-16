package contains_duplicate_ii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_containsDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want bool
	}{
		{
			name: "case 0.1",
			nums: []int{1},
			k:    1,
			want: false,
		},
		{
			name: "case 0.2",
			nums: []int{},
			k:    0,
			want: false,
		},
		{
			name: "case 0.3",
			nums: []int{},
			k:    1,
			want: false,
		},
		{
			name: "case 0.4",
			nums: []int{1},
			k:    0,
			want: false,
		},
		{
			name: "case 0.4",
			nums: []int{1, 2},
			k:    1,
			want: false,
		},
		{
			name: "case 1.1 leetcode",
			nums: []int{1, 2, 3, 1},
			k:    3,
			want: true,
		},
		{
			name: "case 1.2",
			nums: []int{1, 2, 3, 1},
			k:    4,
			want: true,
		},
		{
			name: "case 2 leetcode",
			nums: []int{1, 0, 1, 1},
			k:    1,
			want: true,
		},
		{
			name: "case 3 leetcode",
			nums: []int{1, 2, 3, 1, 2, 3},
			k:    2,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := containsNearbyDuplicate(tt.nums, tt.k)
			assert.Equal(t, tt.want, got)
		})
	}
}
