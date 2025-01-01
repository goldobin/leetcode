package longest_palindromic_substring_recursive

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "case 01",
			s:    "babad",
			want: "bab",
		},
		{
			name: "case 02",
			s:    "cbbd",
			want: "bb",
		},
		{
			name: "case 03",
			s:    "a",
			want: "a",
		},
		{
			name: "case 04",
			s:    "ac",
			want: "a",
		},
		{
			name: "case 05",
			s:    "bb",
			want: "bb",
		},
		{
			name: "case 06",
			s:    "ccc",
			want: "ccc",
		},
		{
			name: "case 07",
			s:    "aaaa",
			want: "aaaa",
		},
		{
			name: "case 08",
			s:    "abcda",
			want: "a",
		},
		{
			name: "case 15",
			s:    "abcdbacdcfgf",
			want: "cdc",
		},
		{
			name: "case 16",
			s:    "kabcbafg",
			want: "abcba",
		},
		{
			name: "case 17",
			s:    "kbbf",
			want: "bb",
		},
		{
			name: "case 18",
			s:    "kbbbbf",
			want: "bbbb",
		},
		{
			name: "case 19",
			s:    "abb",
			want: "bb",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestPalindrome(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
