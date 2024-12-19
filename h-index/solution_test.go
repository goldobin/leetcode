package h_index

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hIndex(t *testing.T) {
	tests := []struct {
		name      string
		citations []int
		want      int
	}{
		{
			name:      "case 1.1 leetcode",
			citations: []int{3, 0, 6, 1, 5},
			want:      3,
		},
		{
			name:      "case 1.2 leetcode",
			citations: []int{1, 3, 1},
			want:      1,
		},
		{
			name:      "case 1.3 leetcode",
			citations: []int{1, 3, 1},
			want:      1,
		},
		{
			name:      "case 1.4 leetcode",
			citations: []int{100},
			want:      1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hIndex(tt.citations)
			assert.Equal(t, tt.want, got)
		})
	}
}
