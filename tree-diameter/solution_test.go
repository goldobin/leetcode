package tree_diameter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_treeDiameter(t *testing.T) {
	tests := []struct {
		name  string
		edges [][]int
		want  int
	}{
		{
			name: "case 0.1 empty edges ",
			want: 0,
		},
		{
			name:  "case 0.2 one edge",
			edges: [][]int{{0, 1}},
			want:  1,
		},
		{
			name:  "case 1.1 leetcode",
			edges: [][]int{{0, 1}, {0, 2}},
			want:  2,
		},
		{
			name: "case 1.2 leetcode",
			edges: [][]int{
				{0, 1},
				{1, 2},
				{2, 3},
				{1, 4},
				{4, 5},
			},
			want: 4,
		},
		{
			name: "case 2.1",
			edges: [][]int{
				{1, 4},
				{2, 4},
				{3, 4},
			},
			want: 2,
		},
		{
			name: "case 2.2",
			edges: [][]int{
				{1, 2},
				{2, 3},
				{3, 4},
				{3, 5},
				{2, 6},
			},
			want: 3,
		},
		{
			name: "case 2.3",
			edges: [][]int{
				{1, 2},
				{2, 3},
				{3, 4},
				{4, 5},
				{5, 6},
				{6, 7},
				{5, 8},
				{2, 9},
			},
			want: 6,
		},
		{
			name: "case 2.2",
			edges: [][]int{
				{0, 1},
				{1, 2},
				{2, 3},
				{1, 4},
				{4, 5},
				{4, 6},
				{6, 7},
				{5, 8},
				{5, 9},
				{8, 10},
			},
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := treeDiameter(tt.edges)
			assert.Equal(t, tt.want, got)
		})
	}
}
