package surrounded_regions

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_solve(t *testing.T) {
	tests := []struct {
		name  string
		board [][]byte
		want  [][]byte
	}{
		{
			name: "case 0.1 empty board",
		},
		{
			name:  "case 0.1.1 1x1 element board",
			board: [][]byte{{'X'}},
			want:  [][]byte{{'X'}},
		},
		{
			name:  "case 0.1.2 1x1 element board",
			board: [][]byte{{'O'}},
			want:  [][]byte{{'O'}},
		},
		{
			name: "case 0.2.1 2x2 element board",
			board: [][]byte{
				{'X', 'X'},
				{'X', 'X'},
			},
			want: [][]byte{
				{'X', 'X'},
				{'X', 'X'},
			},
		},
		{
			name: "case 0.2.2 2x2 element board",
			board: [][]byte{
				{'O', 'O'},
				{'O', 'O'},
			},
			want: [][]byte{
				{'O', 'O'},
				{'O', 'O'},
			},
		},
		{
			name: "case 0.2.3 2x2 element board",
			board: [][]byte{
				{'X', 'O'},
				{'X', 'X'},
			},
			want: [][]byte{
				{'X', 'O'},
				{'X', 'X'},
			},
		},
		{
			name: "case 0.3.1 2x3 element board",
			board: [][]byte{
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
			},
			want: [][]byte{
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
			},
		},
		{
			name: "case 0.3.2 2x3 element board",
			board: [][]byte{
				{'O', 'O', 'O'},
				{'O', 'O', 'O'},
			},
			want: [][]byte{
				{'O', 'O', 'O'},
				{'O', 'O', 'O'},
			},
		},
		{
			name: "case 0.3.3 2x3 element board",
			board: [][]byte{
				{'O', 'X', 'O'},
				{'O', 'O', 'O'},
			},
			want: [][]byte{
				{'O', 'X', 'O'},
				{'O', 'O', 'O'},
			},
		},
		{
			name: "case 0.3.4 2x3 element board",
			board: [][]byte{
				{'X', 'O', 'X'},
				{'X', 'X', 'X'},
			},
			want: [][]byte{
				{'X', 'O', 'X'},
				{'X', 'X', 'X'},
			},
		},
		{
			name: "case 1.1.1 3x3 element board",
			board: [][]byte{
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
			},
			want: [][]byte{
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
			},
		},
		{
			name: "case 1.1.2 3x3 element board",
			board: [][]byte{
				{'X', 'O', 'X'},
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
			},
			want: [][]byte{
				{'X', 'O', 'X'},
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
			},
		},
		{
			name: "case 1.1.3 3x3 element board",
			board: [][]byte{
				{'X', 'O', 'X'},
				{'X', 'X', 'X'},
				{'X', 'O', 'X'},
			},
			want: [][]byte{
				{'X', 'O', 'X'},
				{'X', 'X', 'X'},
				{'X', 'O', 'X'},
			},
		},
		{
			name: "case 1.1.4 3x3 element board",
			board: [][]byte{
				{'X', 'X', 'X'},
				{'O', 'X', 'O'},
				{'X', 'X', 'X'},
			},
			want: [][]byte{
				{'X', 'X', 'X'},
				{'O', 'X', 'O'},
				{'X', 'X', 'X'},
			},
		},
		{
			name: "case 1.1.5 3x3 element board",
			board: [][]byte{
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
				{'X', 'X', 'O'},
			},
			want: [][]byte{
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
				{'X', 'X', 'O'},
			},
		},
		{
			name: "case 1.1.6 3x3 element board",
			board: [][]byte{
				{'X', 'X', 'X'},
				{'X', 'O', 'X'},
				{'X', 'X', 'X'},
			},
			want: [][]byte{
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
			},
		},
		{
			name: "case 1.2.1 4x3 element board",
			board: [][]byte{
				{'X', 'O', 'X'},
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
			},
			want: [][]byte{
				{'X', 'O', 'X'},
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
			},
		},
		{
			name: "case 1.2.2 4x3 element board",
			board: [][]byte{
				{'X', 'O', 'X'},
				{'O', 'X', 'O'},
				{'X', 'X', 'X'},
				{'O', 'X', 'O'},
			},
			want: [][]byte{
				{'X', 'O', 'X'},
				{'O', 'X', 'O'},
				{'X', 'X', 'X'},
				{'O', 'X', 'O'},
			},
		},
		{
			name: "case 1.2.3 4x3 element board",
			board: [][]byte{
				{'X', 'X', 'X'},
				{'X', 'O', 'X'},
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
			},
			want: [][]byte{
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
			},
		},
		{
			name: "case 1.2.4 4x3 element board",
			board: [][]byte{
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
				{'X', 'O', 'X'},
				{'X', 'X', 'X'},
			},
			want: [][]byte{
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
			},
		},
		{
			name: "case 1.2.5 4x3 element board",
			board: [][]byte{
				{'X', 'O', 'X'},
				{'X', 'X', 'X'},
				{'X', 'O', 'X'},
				{'X', 'X', 'X'},
			},
			want: [][]byte{
				{'X', 'O', 'X'},
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
			},
		},
		{
			name: "case 1.2.6 4x3 element board",
			board: [][]byte{
				{'X', 'X', 'X'},
				{'X', 'O', 'X'},
				{'X', 'O', 'X'},
				{'X', 'X', 'X'},
			},
			want: [][]byte{
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
			},
		},
		{
			name: "case 1.2.7 4x3 element board",
			board: [][]byte{
				{'O', 'O', 'O'},
				{'O', 'O', 'O'},
				{'O', 'O', 'O'},
				{'O', 'O', 'O'},
			},
			want: [][]byte{
				{'O', 'O', 'O'},
				{'O', 'O', 'O'},
				{'O', 'O', 'O'},
				{'O', 'O', 'O'},
			},
		},
		{
			name: "case 1.3.1 4x4 element board",
			board: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
			},
			want: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
			},
		},
		{
			name: "case 1.3.2 4x4 element board",
			board: [][]byte{
				{'O', 'O', 'O', 'O'},
				{'O', 'O', 'O', 'O'},
				{'O', 'O', 'O', 'O'},
				{'O', 'O', 'O', 'O'},
			},
			want: [][]byte{
				{'O', 'O', 'O', 'O'},
				{'O', 'O', 'O', 'O'},
				{'O', 'O', 'O', 'O'},
				{'O', 'O', 'O', 'O'},
			},
		},
		{
			name: "case 1.3.3 4x4 element board",
			board: [][]byte{
				{'O', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'X'},
				{'X', 'O', 'O', 'X'},
				{'X', 'X', 'X', 'X'},
			},
			want: [][]byte{
				{'O', 'X', 'O', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
			},
		},
		{
			name: "case 1.4.1 5x5 element board",
			board: [][]byte{
				{'X', 'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X', 'X'},
			},
			want: [][]byte{
				{'X', 'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X', 'X'},
			},
		},
		{
			name: "case 1.4.2 5x5 element board",
			board: [][]byte{
				{'X', 'X', 'X', 'X', 'X'},
				{'X', 'O', 'X', 'O', 'X'},
				{'X', 'O', 'O', 'X', 'X'},
				{'X', 'X', 'X', 'X', 'O'},
			},
			want: [][]byte{
				{'X', 'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X', 'O'},
			},
		},
		{
			name: "case 1.4.3 5x5 element board",
			board: [][]byte{
				{'X', 'X', 'X', 'X', 'X'},
				{'X', 'O', 'X', 'O', 'X'},
				{'X', 'O', 'O', 'X', 'X'},
				{'X', 'X', 'O', 'X', 'O'},
			},
			want: [][]byte{
				{'X', 'X', 'X', 'X', 'X'},
				{'X', 'O', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'X', 'X'},
				{'X', 'X', 'O', 'X', 'O'},
			},
		},
		{
			name: "case 2.1 leetcode",
			board: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'X'},
				{'X', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'X'},
			},
			want: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'X', 'X'},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := clone2(tt.board)
			solve(got)
			assert.Equal(t, tt.want, got)
		})
	}
}

func clone2(src [][]byte) [][]byte {
	if src == nil {
		return nil
	}

	result := make([][]byte, len(src))
	for i, row := range src {
		result[i] = slices.Clone(row)
	}
	return result
}

func Benchmark_solve(b *testing.B) {
	inputs := [][][]byte{
		{
			{'X', 'X', 'X', 'X', 'X'},
			{'X', 'O', 'X', 'O', 'X'},
			{'X', 'O', 'O', 'X', 'X'},
			{'X', 'X', 'X', 'X', 'O'},
		},
		{
			{'X', 'X', 'X', 'X', 'X'},
			{'X', 'O', 'X', 'O', 'X'},
			{'X', 'O', 'O', 'X', 'X'},
			{'X', 'X', 'O', 'X', 'O'},
		},
	}

	for b.Loop() {
		for _, in := range inputs {
			solve(in)
		}
	}
}
