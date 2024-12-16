package climb_stairs_recursive

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_climbStirs(t *testing.T) {
	tests := []struct {
		name string
		in   int
		want int
	}{
		{
			name: "case 1 leetcode",
			in:   2,
			want: 2,
		},
		{
			name: "case 2 leetcode",
			in:   3,
			want: 3,
		},
		{
			name: "case 3",
			in:   5,
			want: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := climbStairs(tt.in)
			assert.Equal(t, tt.want, got)
		})
	}
}
