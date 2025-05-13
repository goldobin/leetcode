package powx_n

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_myPow(t *testing.T) {
	tests := []struct {
		name string
		x    float64
		n    int
		want float64
	}{
		{
			name: "0.1 0^1",
			x:    0,
			n:    1,
			want: 0,
		},
		{
			name: "0.2 1^1",
			x:    1,
			n:    1,
			want: 1,
		},
		{
			name: "0.3 -1^1",
			x:    -1,
			n:    1,
			want: -1,
		},
		{
			name: "0.4 -2^1",
			x:    -2,
			n:    1,
			want: -2,
		},
		{
			name: "1.1 0^0",
			x:    0,
			n:    0,
			want: 1,
		},
		{
			name: "1.2 1^0",
			x:    1,
			n:    0,
			want: 1,
		},
		{
			name: "1.3 2^0",
			x:    2,
			n:    0,
			want: 1,
		},
		{
			name: "1.4 2^0",
			x:    -2,
			n:    0,
			want: 1,
		},
		{
			name: "2.1 -2^2",
			x:    2,
			n:    2,
			want: 4,
		},
		{
			name: "2.2 -2^2",
			x:    -2,
			n:    2,
			want: 4,
		},
		{
			name: "3.1 leetcode",
			x:    2,
			n:    10,
			want: 1024,
		},
		{
			name: "3.2 leetcode",
			x:    2.1,
			n:    3,
			want: 9.261,
		},
		{
			name: "3.3 leetcode",
			x:    2,
			n:    -2,
			want: 0.25,
		},
		{
			name: "3.4 leetcode",
			x:    34.00515,
			n:    -3,
			want: 3e-05,
		},
		{
			name: "4.1 big powers",
			x:    2,
			n:    -int(math.Pow(2, 31)),
			want: math.Pow(2, -math.Pow(2, 31)),
		},
		{
			name: "4.2 big powers",
			x:    2,
			n:    int(math.Pow(2, 31) - 1),
			want: math.Pow(2, math.Pow(2, 31)-1),
		},
		{
			name: "4.3 big powers",
			x:    -2,
			n:    -int(math.Pow(2, 31)),
			want: math.Pow(-2, -math.Pow(2, 31)),
		},
		{
			name: "4.4 big powers",
			x:    -2,
			n:    int(math.Pow(2, 31) - 1),
			want: math.Pow(-2, math.Pow(2, 31)-1),
		},
		{
			name: "4.5 big powers",
			x:    -99.999,
			n:    -int(math.Pow(2, 31)),
			want: math.Pow(-99.999, -math.Pow(2, 31)),
		},
		{
			name: "4.6 big powers",
			x:    -99.999,
			n:    int(math.Pow(2, 31) - 1),
			want: math.Pow(-99.999, math.Pow(2, 31)-1),
		},
		{
			name: "4.7 big powers",
			x:    99.999,
			n:    -int(math.Pow(2, 31)),
			want: math.Pow(99.999, -math.Pow(2, 31)),
		},
		{
			name: "4.8 big powers",
			x:    99.999,
			n:    int(math.Pow(2, 31) - 1),
			want: math.Pow(99.999, math.Pow(2, 31)-1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := myPow(tt.x, tt.n)
			assert.InDelta(t, tt.want, got, 0.0001)
		})
	}
}
