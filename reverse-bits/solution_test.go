package reverse_bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseBits(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "case 0.1 zero",
			want: 0,
		},
		{
			name: "case 0.2 one",
			n:    1,
			want: 0b10000000_00000000_00000000_00000000,
		},
		{
			name: "case 0.3 reverse one",
			n:    0b10000000_00000000_00000000_00000000,
			want: 0b1,
		},
		{
			name: "case 0.4 all ones",
			n:    0b11111111_11111111_11111111_11111111,
			want: 0b11111111_11111111_11111111_11111111,
		},
		{
			name: "case 1.1 leetcode",
			n:    43261596,
			want: 964176192,
		},
		{
			name: "case 1.2 leetcode",
			n:    2147483644,
			want: 1073741822,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseBits(tt.n)
			assert.Equal(t, tt.want, got)
		})
	}
}
