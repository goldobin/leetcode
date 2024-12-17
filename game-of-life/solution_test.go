package game_of_life

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_gameOfLife(t *testing.T) {
	tests := []struct {
		name  string
		board [][]int
		want  [][]int
	}{
		{
			name: "case 1 leetcode",
			board: [][]int{
				{0, 1, 0},
				{0, 0, 1},
				{1, 1, 1},
				{0, 0, 0},
			},
			want: [][]int{
				{0, 0, 0},
				{1, 0, 1},
				{0, 1, 1},
				{0, 1, 0},
			},
		},
		{
			name: "case 2 leetcode",
			board: [][]int{
				{1, 1},
				{1, 0},
			},
			want: [][]int{
				{1, 1},
				{1, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := slices.Clone(tt.board)
			gameOfLife(got)

			assert.Equal(t, tt.want, got)
		})
	}
}
