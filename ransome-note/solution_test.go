package ransome_note

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_canConstruct(t *testing.T) {
	tests := []struct {
		name       string
		ransomNote string
		magazine   string
		want       bool
	}{
		{
			name: "case 0.1 empty note and magazine",
			want: true,
		},
		{
			name:       "case 0.2 empty magazine",
			ransomNote: "a",
			want:       false,
		},
		{
			name:       "case 0.3 one letter note and magazine, same letter",
			ransomNote: "a",
			magazine:   "a",
			want:       true,
		},
		{
			name:       "case 0.4 one letter note and magazine, different letters",
			ransomNote: "a",
			magazine:   "b",
			want:       false,
		},
		{
			name:       "case 0.4 one letter note and muti letter magazine, contains letters",
			ransomNote: "a",
			magazine:   "ab",
			want:       true,
		},
		{
			name:       "case 0.5 one letter note and muti letter magazine, doesnt contain letters",
			ransomNote: "a",
			magazine:   "bc",
			want:       false,
		},
		{
			name:       "case 0.6 two letter note and muti letter magazine, contains letters",
			ransomNote: "ab",
			magazine:   "abc",
			want:       true,
		},
		{
			name:       "case 0.7 two letter note and muti letter magazine, contains some letters",
			ransomNote: "ab",
			magazine:   "bc",
			want:       false,
		},
		{
			name:       "case 0.8 two letter note and muti letter magazine, contains all letters",
			ransomNote: "aa",
			magazine:   "aabb",
			want:       true,
		},
		{
			name:       "case 0.9 two letter note and muti letter magazine, contains all letters but not enough count",
			ransomNote: "aa",
			magazine:   "abb",
			want:       false,
		},
		{
			name: "case 1.1 leetcode",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canConstruct(tt.ransomNote, tt.magazine)
			assert.Equal(t, tt.want, got)
		})
	}
}
