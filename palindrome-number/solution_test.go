package palindrome_number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want bool
	}{
		{
			name: "case 0.1 zero",
			x:    0,
			want: true,
		},
		{
			name: "case 0.2 minus one",
			x:    -1,
			want: false,
		},
		{
			name: "case 0.3 minus 101",
			x:    -101,
			want: false,
		},
		{
			name: "case 0.4 one",
			x:    1,
			want: true,
		},
		{
			name: "case 0.5 eleven",
			x:    11,
			want: true,
		},
		{
			name: "case 0.6 twenty",
			x:    20,
			want: false,
		},
		{
			name: "case 0.7 twenty two",
			x:    22,
			want: true,
		},
		{
			name: "case 0.8 thirty two",
			x:    32,
			want: false,
		},
		{
			name: "case 0.9 one zero one",
			x:    101,
			want: true,
		},
		{
			name: "case 0.10 five two five",
			x:    525,
			want: true,
		},
		{
			name: "case 0.11 five zero five",
			x:    505,
			want: true,
		},
		{
			name: "case 0.12 five zero one",
			x:    501,
			want: false,
		},
		{
			name: "case 1.1 leetcode",
			x:    121,
			want: true,
		},
		{
			name: "case 1.2 leetcode",
			x:    -121,
			want: false,
		},
		{
			name: "case 1.2 leetcode",
			x:    10,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPalindrome(tt.x)
			assert.Equal(t, tt.want, got)
		})
	}
}
