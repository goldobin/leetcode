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
		want result
	}{
		{
			name: "case 01.1 empty strings",
			s:    "",
			t:    "",
			want: result{
				cost:     0,
				editPath: []int{},
			},
		},
		{
			name: "case 01.2.1 one empty string",
			s:    "",
			t:    "a",
			want: result{
				cost:     1,
				editPath: []int{insertOp},
			},
		},
		{
			name: "case 01.2.2 one empty string",
			s:    "a",
			t:    "",
			want: result{
				cost:     1,
				editPath: []int{deleteOp},
			},
		},
		{
			name: "case 01.3.1 one empty string",
			s:    "",
			t:    "ab",
			want: result{
				cost:     2,
				editPath: []int{insertOp, insertOp},
			},
		},
		{
			name: "case 01.3.2 one empty string",
			s:    "ab",
			t:    "",
			want: result{
				cost:     2,
				editPath: []int{deleteOp, deleteOp},
			},
		},
		{
			name: "case 01.3.1 one empty string",
			s:    "",
			t:    "abc",
			want: result{
				cost:     3,
				editPath: []int{insertOp, insertOp, insertOp},
			},
		},
		{
			name: "case 01.3.2 one empty string",
			s:    "abc",
			t:    "",
			want: result{
				cost:     3,
				editPath: []int{deleteOp, deleteOp, deleteOp},
			},
		},
		{
			name: "case 02.1 equal non empty strings",
			s:    "a",
			t:    "a",
			want: result{
				cost:     0,
				editPath: []int{matchOp},
			},
		},
		{
			name: "case 02.2 non equal non empty strings",
			s:    "a",
			t:    "b",
			want: result{
				cost:     1,
				editPath: []int{matchOp},
			},
		},
		{
			name: "case 02.3 non equal non empty strings",
			s:    "a",
			t:    "ab",
			want: result{
				cost:     1,
				editPath: []int{matchOp, insertOp},
			},
		},
		{
			name: "case 02.4 non equal non empty strings",
			s:    "abc",
			t:    "ac",
			want: result{
				cost:     1,
				editPath: []int{matchOp, deleteOp, matchOp},
			},
		},
		{
			name: "case 02.5 non equal non empty strings",
			s:    "abc",
			t:    "ad",
			want: result{
				cost:     2,
				editPath: []int{matchOp, deleteOp, matchOp},
			},
		},
		{
			name: "case 02.6 non equal non empty strings",
			s:    "abc",
			t:    "dab",
			want: result{
				cost: 2,
				editPath: []int{
					insertOp, matchOp, matchOp, deleteOp,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := stringCompare(tt.s, tt.t)
			assert.Equal(t, tt.want, got)
		})
	}
}
