package number_of_1_bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hammingWeight(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "case 0.1 zero",
			n:    0,
			want: 0,
		},
		{
			name: "case 1.1 one",
			n:    1,
			want: 1,
		},
		{
			name: "case 1.2 one",
			n:    0b10,
			want: 1,
		},
		{
			name: "case 1.3 2",
			n:    0b11,
			want: 2,
		},
		{
			name: "case 1.3 2",
			n:    0xffffffff,
			want: 32,
		},
		{
			name: "case 2.1 leetcode",
			n:    128,
			want: 1,
		},
		{
			name: "case 2.2 leetcode",
			n:    2147483645,
			want: 30,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hammingWeight(tt.n)
			assert.Equal(t, tt.want, got)
		})
	}
}
