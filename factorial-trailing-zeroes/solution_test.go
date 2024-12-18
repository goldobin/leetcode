package factorial_trailing_zeroes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_trailingZeroes(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "case 1.1 leetcode",
			n:    3,
			want: 0,
		},
		{
			name: "case 1.2 leetcode",
			n:    5,
			want: 1,
		},
		{
			name: "case 1.3 leetcode",
			n:    10,
			want: 2,
		},
		{
			name: "case 2.1",
			n:    11,
			want: 2,
		},
		{
			name: "case 2.2",
			n:    14,
			want: 2,
		},
		{
			name: "case 2.3",
			n:    15,
			want: 3,
		},
		{
			name: "case 2.4",
			n:    17,
			want: 3,
		},
		{
			name: "case 2.5",
			n:    20,
			want: 4,
		},
		{
			name: "case 2.6",
			n:    30,
			want: 7,
		},
		{
			name: "case 2.7 ",
			n:    40,
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := trailingZeroes(tt.n)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Benchmark_trailingZeroes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		trailingZeroes(1000)
	}
}
