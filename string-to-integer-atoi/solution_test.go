package string_to_integer_atoi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_myAtoi(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "case 0.1.1",
			s:    "",
			want: 0,
		},
		{
			name: "case 0.1.2",
			s:    "0",
			want: 0,
		},
		{
			name: "case 0.1.3",
			s:    "1",
			want: 1,
		},
		{
			name: "case 0.1.4",
			s:    "-1",
			want: -1,
		},
		{
			name: "case 0.1.5",
			s:    "2",
			want: 2,
		},
		{
			name: "case 0.1.5",
			s:    "10",
			want: 10,
		},
		{
			name: "case 0.1.5",
			s:    " -10",
			want: -10,
		},
		{
			name: "case 0.2 leetcode",
			s:    "0-1",
			want: 0,
		},
		{
			name: "case 0.2",
			s:    "01-1",
			want: 1,
		},
		{
			name: "case 0.3",
			s:    "09-1",
			want: 9,
		},
		{
			name: "case 0.4 leetcode",
			s:    "words and 987",
			want: 0,
		},
		{
			name: "case 0.5 leetcode",
			s:    "words and 987",
			want: 0,
		},
		{
			name: "case 1.1 leetcode",
			s:    "42",
			want: 42,
		},
		{
			name: "case 1.2 leetcode",
			s:    " -042",
			want: -42,
		},
		{
			name: "case 1.3.1 leetcode",
			s:    "1337c0d3",
			want: 1337,
		},
		{
			name: "case 1.3.2",
			s:    "   1337c0d3",
			want: 1337,
		},
		{
			name: "case 1.3.3",
			s:    "  001337c0d3",
			want: 1337,
		},
		{
			name: "case 1.3.4",
			s:    "  00-1337c0d3",
			want: 0,
		},
		{
			name: "case 1.3.5",
			s:    "  -001337c0d3",
			want: -1337,
		},
		{
			name: "case 1.3.6",
			s:    "-001337c0d3",
			want: -1337,
		},
		{
			name: "case 1.3.7",
			s:    "0000001337c0d3",
			want: 1337,
		},
		{
			name: "case 1.3.8",
			s:    "00000010337c0d3",
			want: 10337,
		},
		{
			name: "case 1.3.9",
			s:    "+00000010337c0d3",
			want: 10337,
		},
		{
			name: "case 1.3.10",
			s:    "   +1337c0d3",
			want: 1337,
		},
		{
			name: "case 1.3.11",
			s:    "   +  1337c0d3",
			want: 0,
		},
		{
			name: "case 1.3.12",
			s:    "   -  1337c0d3",
			want: 0,
		},
		{
			name: "case 1.3.13",
			s:    "   -  13378",
			want: 0,
		},
		{
			name: "case 1.3.14",
			s:    "   -  1337+8",
			want: 0,
		},
		{
			name: "case 1.4.1 leetcode",
			s:    "-91283472332",
			want: -2147483648,
		},
		{
			name: "case 1.4.2",
			s:    "-2147483648",
			want: -2147483648,
		},
		{
			name: "case 1.4.3",
			s:    "-2147483649",
			want: -2147483648,
		},
		{
			name: "case 1.4.4",
			s:    "-2147483647",
			want: -2147483647,
		},
		{
			name: "case 1.5.1",
			s:    "2147483647",
			want: 2147483647,
		},
		{
			name: "case 1.5.2",
			s:    "2147483648",
			want: 2147483647,
		},
		{
			name: "case 1.5.3",
			s:    "02147483648",
			want: 2147483647,
		},
		{
			name: "case 1.5.4",
			s:    "2147483646",
			want: 2147483646,
		},
		{
			name: "case 1.6 leetcode",
			s:    "  +  413",
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := myAtoi(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
