package most_profitable_path_in_a_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mostProfitablePath(t *testing.T) {
	tests := []struct {
		name         string
		edges        [][]int
		bobStartNode int
		rewards      []int
		want         int
	}{
		{
			name:         "case 1.1 leetcode",
			edges:        [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}},
			bobStartNode: 3,
			rewards:      []int{-2, 4, 2, -4, 6},
			want:         6,
		},
		{
			name:         "case 1.2 leetcode",
			edges:        [][]int{{0, 1}},
			bobStartNode: 1,
			rewards:      []int{-7280, 2350},
			want:         -7280,
		},
		{
			name:         "case 1.3 leetcode",
			edges:        [][]int{{0, 1}, {1, 2}, {2, 3}},
			bobStartNode: 3,
			rewards:      []int{-5644, -6018, 1188, -8502},
			want:         -11662,
		},
		{
			name:         "case 1.4 leetcode",
			edges:        [][]int{{0, 2}, {0, 4}, {1, 3}, {1, 2}},
			bobStartNode: 1,
			rewards:      []int{3958, -9854, -8334, -9388, 3410},
			want:         7368,
		},
		{
			name:         "case 2.1",
			edges:        [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}, {2, 6}},
			bobStartNode: 3,
			rewards:      []int{10, 20, 30, 40, 50, 60, 70},
			want:         10 + 30 + 70,
		},
		{
			name:         "case 2.2",
			edges:        [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}, {2, 6}},
			bobStartNode: 6,
			rewards:      []int{10, 20, 30, 40, 50, 60, 70},
			want:         10 + 30/2 + 60,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mostProfitablePath(tt.edges, tt.bobStartNode, tt.rewards)
			assert.Equal(t, tt.want, got)
		})
	}
}
