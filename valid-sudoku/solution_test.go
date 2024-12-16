package valid_sudoku

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isValidSudoku(t *testing.T) {
	tests := []struct {
		name  string
		board [][]byte
		want  bool
	}{
		{
			name: "case 1.1 leetcode",
			board: [][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			want: true,
		},
		{
			name: "case 1.2 leetcode",
			board: [][]byte{
				{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValidSudoku(tt.board)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_boxIndex(t *testing.T) {
	tests := []struct {
		name string
		i, j int
		want int
	}{
		{
			name: "case 0,0",
			i:    0,
			j:    0,
			want: 0,
		},
		{
			name: "case 0,2",
			i:    0,
			j:    2,
			want: 0,
		},
		{
			name: "case 0,3",
			i:    0,
			j:    3,
			want: 1,
		},
		{
			name: "case 0,8",
			i:    0,
			j:    8,
			want: 2,
		},
		{
			name: "case 1,0",
			i:    1,
			j:    0,
			want: 0,
		},
		{
			name: "case 3,1",
			i:    3,
			j:    1,
			want: 3,
		},
		{
			name: "case 4,4",
			i:    4,
			j:    4,
			want: 4,
		},
		{
			name: "case 7,7",
			i:    7,
			j:    7,
			want: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := boxIndex(tt.i, tt.j)
			assert.Equal(t, tt.want, got)
		})
	}
}
