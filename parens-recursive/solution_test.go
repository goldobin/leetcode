package parens_recursive

import (
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	tests := []struct {
		name string
		in   int
		want []string
	}{
		{
			name: "case 0",
			in:   0,
			want: []string{},
		},
		{
			name: "case 1 leetcode",
			in:   1,
			want: []string{
				"()",
			},
		},
		{
			name: "case 2 leetcode",
			in:   2,
			want: []string{
				"(())",
				"()()",
			},
		},
		{
			name: "case 3 leetcode",
			in:   3,
			want: []string{
				"((()))",
				"(()())",
				"(())()",
				"()(())",
				"()()()",
			},
		},
		{
			name: "case 3 leetcode",
			in:   4,
			want: []string{
				"(((())))",
				"((()()))",
				"((())())",
				"((()))()",
				"(()(()))",
				"(()()())",
				"(()())()",
				"(())(())",
				"(())()()",
				"()((()))",
				"()(()())",
				"()(())()",
				"()()(())",
				"()()()()",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				got        = generateParenthesis(tt.in)
				wantSorted = slices.Clone(tt.want)
				gotSorted  = slices.Clone(got)
			)

			slices.Sort(wantSorted)
			slices.Sort(gotSorted)
			assert.Equal(t, wantSorted, gotSorted)
		})
	}
}
