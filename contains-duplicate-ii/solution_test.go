package contains_duplicate_ii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_containsDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want bool
	}{
		{
			name: "case 1 leetcode",
			nums: []int{1, 2, 3, 1},
			k:    3,
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
