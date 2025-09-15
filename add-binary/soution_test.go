package add_binary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_addBinary(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
		want string
	}{
		{
			name: "test 0.1 empty operands",
		},
		{
			name: "test 0.2.1 one operand is empty",
			a:    "1",
			want: "1",
		},
		{
			name: "test 0.2.2 one operand is empty",
			a:    "10",
			want: "10",
		},
		{
			name: "test 1.1 1 and 1",
			a:    "1",
			b:    "1",
			want: "10",
		},
		{
			name: "test 1.2 0 and 0",
			a:    "0",
			b:    "0",
			want: "0",
		},
		{
			name: "test 1.3 0 and 1",
			a:    "0",
			b:    "1",
			want: "1",
		},
		{
			name: "test 1.4 10 and 1",
			a:    "10",
			b:    "1",
			want: "11",
		},
		{
			name: "test 1.5 10 and 0",
			a:    "10",
			b:    "0",
			want: "10",
		},
		{
			name: "test 2.1 leetcode",
			a:    "11",
			b:    "1",
			want: "100",
		},
		{
			name: "test 2.3 leetcode",
			a:    "1010",
			b:    "1011",
			want: "10101",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1 := addBinary(tt.a, tt.b)
			got2 := addBinary(tt.b, tt.a)
			assert.Equal(t, tt.want, got1)
			assert.Equal(t, tt.want, got2)
		})
	}
}
