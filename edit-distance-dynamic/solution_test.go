package edit_distance_dynamic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_stringCompare(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want int
	}{
		{
			name: "case 1.1 empty strings",
			s:    "",
			t:    "",
			want: 0,
		},
		{
			name: "case 1.2.1 one empty string",
			s:    "",
			t:    "a",
			want: 1,
		},
		{
			name: "case 1.2.2 one empty string",
			s:    "a",
			t:    "",
			want: 1,
		},
		{
			name: "case 1.3.1 one empty string",
			s:    "",
			t:    "ab",
			want: 2,
		},
		{
			name: "case 1.3.2 one empty string",
			s:    "ab",
			t:    "",
			want: 2,
		},
		{
			name: "case 1.3.1 one empty string",
			s:    "",
			t:    "abc",
			want: 3,
		},
		{
			name: "case 1.3.2 one empty string",
			s:    "abc",
			t:    "",
			want: 3,
		},
		{
			name: "case 2.1 equal non empty strings",
			s:    "a",
			t:    "a",
			want: 0,
		},
		{
			name: "case 2.2 non equal non empty strings",
			s:    "a",
			t:    "b",
			want: 1,
		},
		{
			name: "case 2.3 non equal non empty strings",
			s:    "a",
			t:    "ab",
			want: 1,
		},
		{
			name: "case 2.4 non equal non empty strings",
			s:    "abc",
			t:    "ac",
			want: 1,
		},
		{
			name: "case 2.5 non equal non empty strings",
			s:    "abc",
			t:    "ad",
			want: 2,
		},
		{
			name: "case 2.6 non equal non empty strings",
			s:    "abc",
			t:    "dab",
			want: 2,
		},
		{
			name: "case 3.1 leetcode",
			s:    "horse",
			t:    "ros",
			want: 3,
		},
		{
			name: "case 3.2 leetcode",
			s:    "intention",
			t:    "execution",
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minDistance(tt.s, tt.t)
			assert.Equal(t, tt.want, got)
		})
	}
}
