package container_with_most_water

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxArea(t *testing.T) {
	tests := []struct {
		name   string
		height []int
	}{
		{
			name:   "case 0.1",
			height: []int{1},
		},
		{
			name:   "case 0.2",
			height: []int{2, 2, 2, 2, 2},
		},
		{
			name:   "case 1.1 leetcode",
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
		},
		{
			name:   "case 1.2 leetcode",
			height: []int{1, 1},
		},
		{
			name:   "case 1.3 leetcode",
			height: []int{1, 2},
		},
		{
			name:   "case 1.3",
			height: []int{11, 1, 3, 13, 4, 10, 1, 5, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxArea(tt.height)
			gotBf := maxAreaBruteForce(tt.height)
			assert.Equal(t, gotBf, got)
		})
	}
}
