package triangle_recursive

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minimumTotal(t *testing.T) {
	tests := []struct {
		name     string
		triangle [][]int
		want     int
	}{
		{
			name:     "case 0.1",
			triangle: [][]int{},
			want:     0,
		},
		{
			name: "case 1.1 leetcode",
			triangle: [][]int{
				{2},
				{3, 4},
				{6, 5, 7},
				{4, 1, 8, 3},
			},
			want: 11,
		},
		{
			name: "case 1.2 leetcode",
			triangle: [][]int{
				{-10},
			},
			want: -10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimumTotal(tt.triangle)
			assert.Equal(t, tt.want, got)
		})
	}
}
