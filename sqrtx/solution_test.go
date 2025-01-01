package sqrtx

import (
	"strconv"
	"testing"
)

func Test_mySqrt(t *testing.T) {
	tests := []struct {
		x    int
		want int
	}{
		{x: 0, want: 0},
		{x: 1, want: 1},
		{x: 2, want: 1},
		{x: 3, want: 1},
		{x: 4, want: 2},
		{x: 8, want: 2},
		{x: 9, want: 3},
		{x: 15, want: 3},
		{x: 16, want: 4},
		{x: 17, want: 4},
		{x: 26, want: 5},
		{x: 6466, want: 80},
		{x: 64313164, want: 8019},
		{x: 2147395599, want: 46339},
	}

	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.x), func(t *testing.T) {
			if got := mySqrt(tt.x); got != tt.want {
				t.Errorf("mySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
