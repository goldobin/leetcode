package plus_one

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_plusOne(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "test 0.1",
			nums: []int{},
			want: []int{1},
		},
		{
			name: "test 0.2",
			nums: []int{1},
			want: []int{2},
		},
		{
			name: "test 0.3",
			nums: []int{9},
			want: []int{1, 0},
		},
		{
			name: "test 1.1 leetcode",
			nums: []int{1, 2, 3},
			want: []int{1, 2, 4},
		},
		{
			name: "test 1.2 leetcode",
			nums: []int{4, 3, 2, 1},
			want: []int{4, 3, 2, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := plusOne(tt.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}
