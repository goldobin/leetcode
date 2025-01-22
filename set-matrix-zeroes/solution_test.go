package set_matrix_zeroes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_setZeros(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   [][]int
	}{
		{
			name: "case 0.0 nil matrix",
		},
		{
			name:   "case 0.1 0x0 matrix",
			matrix: make([][]int, 0),
			want:   make([][]int, 0),
		},
		{
			name:   "case 0.2 1x0 matrix",
			matrix: make([][]int, 1),
			want:   make([][]int, 1),
		},
		{
			name:   "case 0.3 1x1 zero matrix",
			matrix: [][]int{{0}},
			want:   [][]int{{0}},
		},
		{
			name:   "case 0.4 1x1 one matrix",
			matrix: [][]int{{1}},
			want:   [][]int{{1}},
		},
		{
			name: "case 0.6 2x2 zero matrix",
			matrix: [][]int{
				{0, 0},
				{0, 0},
			},
			want: [][]int{
				{0, 0},
				{0, 0},
			},
		},
		{
			name: "case 0.7 2x2 one matrix",
			matrix: [][]int{
				{1, 1},
				{1, 1},
			},
			want: [][]int{
				{1, 1},
				{1, 1},
			},
		},
		{
			name: "case 0.8.1 2x2 one zero matrix",
			matrix: [][]int{
				{0, 1},
				{1, 1},
			},
			want: [][]int{
				{0, 0},
				{0, 1},
			},
		},
		{
			name: "case 0.8.2 2x2 one zero matrix",
			matrix: [][]int{
				{1, 0},
				{1, 1},
			},
			want: [][]int{
				{0, 0},
				{1, 0},
			},
		},
		{
			name: "case 0.8.3 2x2 one zero matrix",
			matrix: [][]int{
				{1, 1},
				{1, 0},
			},
			want: [][]int{
				{1, 0},
				{0, 0},
			},
		},
		{
			name: "case 0.8.4 2x2 one zero matrix",
			matrix: [][]int{
				{1, 1},
				{0, 1},
			},
			want: [][]int{
				{0, 1},
				{0, 0},
			},
		},
		{
			name: "case 0.9.1 2x2 two zeros matrix",
			matrix: [][]int{
				{1, 0},
				{0, 1},
			},
			want: [][]int{
				{0, 0},
				{0, 0},
			},
		},
		{
			name: "case 0.9.2 2x2 two zeros matrix",
			matrix: [][]int{
				{0, 1},
				{1, 0},
			},
			want: [][]int{
				{0, 0},
				{0, 0},
			},
		},
		{
			name: "case 0.9.3 2x2 two zeros matrix",
			matrix: [][]int{
				{1, 0},
				{1, 0},
			},
			want: [][]int{
				{0, 0},
				{0, 0},
			},
		},
		{
			name: "case 0.9.4 2x2 two zeros matrix",
			matrix: [][]int{
				{0, 1},
				{0, 1},
			},
			want: [][]int{
				{0, 0},
				{0, 0},
			},
		},
		{
			name: "case 1.1 leetcode",
			matrix: [][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			want: [][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
		},
		{
			name: "case 1.2 leetcode",
			matrix: [][]int{
				{0, 1, 2, 0},
				{3, 4, 5, 2},
				{1, 3, 1, 5},
			},
			want: [][]int{
				{0, 0, 0, 0},
				{0, 4, 5, 0},
				{0, 3, 1, 0},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := deepCopy(tt.matrix)
			setZeroes(m)
			assert.Equal(t, tt.want, m)
		})
	}
}

func deepCopy(matrix [][]int) [][]int {
	if matrix == nil {
		return nil
	}

	if len(matrix) == 0 {
		return [][]int{}
	}

	m := make([][]int, len(matrix))
	for i, row := range matrix {
		if row == nil {
			m[i] = nil
			continue
		}

		m[i] = make([]int, len(row))
		copy(m[i], row)
	}
	return m
}
