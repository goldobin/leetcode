package valid_parentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isValid(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "case 0.1",
			input: "",
			want:  true,
		},
		{
			name:  "case 0.2",
			input: "(",
			want:  false,
		},
		{
			name:  "case 0.3",
			input: ")",
			want:  false,
		},
		{
			name:  "case 0.4",
			input: "{",
			want:  false,
		},
		{
			name:  "case 0.5",
			input: "}",
			want:  false,
		},
		{
			name:  "case 0.6",
			input: "[",
			want:  false,
		},
		{
			name:  "case 0.7",
			input: "]",
			want:  false,
		},
		{
			name:  "case 0.8",
			input: "()",
			want:  true,
		},
		{
			name:  "case 0.9",
			input: "[]",
			want:  true,
		},
		{
			name:  "case 0.8",
			input: "{}",
			want:  true,
		},
		{
			name:  "case 1.1",
			input: "{()}",
			want:  true,
		},
		{
			name:  "case 1.2",
			input: "{[]}",
			want:  true,
		},
		{
			name:  "case 1.3",
			input: "{[()]}",
			want:  true,
		},
		{
			name:  "case 1.4",
			input: "{()]}",
			want:  false,
		},
		{
			name:  "case 1.5",
			input: "{[(])}",
			want:  false,
		},
		{
			name:  "case 2.1 leetcode",
			input: "()[]{}",
			want:  true,
		},
		{
			name:  "case 2.2 leetcode",
			input: "(]",
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValid(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
