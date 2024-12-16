package edit_distance_recursive

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
			name: "case 01.1 empty strings",
			s:    "",
			t:    "",
			want: 0,
		},
		{
			name: "case 01.2 one empty string",
			s:    "",
			t:    "a",
			want: 1,
		},
		{
			name: "case 02.1 equal non empty strings",

			s:    "a",
			t:    "a",
			want: 0,
		},
		{
			name: "case 02.2 non equal non empty strings",

			s:    "a",
			t:    "b",
			want: 1,
		},
		{
			name: "case 02.3 non equal non empty strings",

			s:    "a",
			t:    "ab",
			want: 1,
		},
		{
			name: "case 02.4 non equal non empty strings",

			s:    "abc",
			t:    "ac",
			want: 1,
		},
		{
			name: "case 02.5 non equal non empty strings",

			s:    "abc",
			t:    "ad",
			want: 2,
		},
		{
			name: "case 02.6 non equal non empty strings",

			s:    "abc",
			t:    "dab",
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1 := stringCompare(tt.s, tt.t)
			got2 := stringCompare(tt.t, tt.s)
			assert.Equal(t, tt.want, got1)
			assert.Equal(t, tt.want, got2)
		})
	}
}
