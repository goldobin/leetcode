package dot_product_of_two_sparse_vectors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SparseVector_dotProduct(t *testing.T) {
	tests := []struct {
		name string
		v1   []int
		v2   []int
		want int
	}{
		{
			name: "case 0.1 zero-size",
			v1:   []int{},
			v2:   []int{},
			want: 0,
		},
		{
			name: "case 0.2 zeros",
			v1:   []int{0},
			v2:   []int{0},
			want: 0,
		},
		{
			name: "case 0.3 zeros",
			v1:   []int{0},
			v2:   []int{0},
			want: 0,
		},
		{
			name: "case 1.1 leetcode",
			v1:   []int{1, 0, 0, 2, 3},
			v2:   []int{0, 3, 0, 4, 0},
			want: 8,
		},
		{
			name: "case 1.2 leetcode",
			v1:   []int{0, 1, 0, 0, 0},
			v2:   []int{0, 0, 0, 0, 2},
			want: 0,
		},
		{
			name: "case 1.3 leetcode",
			v1:   []int{0, 1, 0, 0, 2, 0, 0},
			v2:   []int{1, 0, 0, 0, 3, 0, 4},
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				v1   = Constructor(tt.v1)
				v2   = Constructor(tt.v2)
				got1 = v1.dotProduct(v2)
				got2 = v2.dotProduct(v1)
			)

			assert.Equal(t, tt.want, got1, tt.name)
			assert.Equal(t, tt.want, got2, tt.name)
		})
	}
}

func Test_SparseVector_dotProduct_differentSizes(t *testing.T) {
	tests := []struct {
		name string
		v1   []int
		v2   []int
	}{
		{
			name: "case 0.2.1 different sizes",
			v1:   []int{},
			v2:   []int{1},
		},
		{
			name: "case 0.2.2 different sizes",
			v1:   []int{1},
			v2:   []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				v1 = Constructor(tt.v1)
				v2 = Constructor(tt.v2)
			)

			assert.Panics(t, func() {
				_ = v1.dotProduct(v2)
			})
			assert.Panics(t, func() {
				_ = v2.dotProduct(v1)
			})
		})
	}
}
