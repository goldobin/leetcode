package subarray_sum_equals_k

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_subarraySum(t *testing.T) {
	tests := []struct {
		name  string
		inArr []int
		inK   int
		want  int
	}{
		{
			name:  "case 01 leetcode",
			inArr: []int{1, 1, 1},
			inK:   2,
			want:  2,
		},
		{
			name:  "case 01.1 mod",
			inArr: []int{1, 1, -1, 1},
			inK:   2,
			want:  2,
		},
		{
			name:  "case 02 leetcode",
			inArr: []int{1, 2, 3},
			inK:   3,
			want:  2,
		},
		{
			name:  "case 03",
			inArr: []int{1, 1, 3},
			inK:   3,
			want:  1,
		},
		{
			name:  "case 03.1",
			inArr: []int{1, -1, 1, 1, 3},
			inK:   3,
			want:  1,
		},
		{
			name:  "case 03.2",
			inArr: []int{1, 1, -1, 1, -1, 1, 1, 3},
			inK:   3,
			want:  2,
		},
		{
			name:  "case 04 leetcode",
			inArr: []int{1, -1, 0},
			inK:   0,
			want:  3,
		},
		{
			name:  "case 05 leetcode",
			inArr: []int{1},
			inK:   0,
			want:  0,
		},
		{
			name:  "case 06 leetcode",
			inArr: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			inK:   0,
			want:  55,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := subarraySum(tt.inArr, tt.inK)
			assert.Equal(t, tt.want, got)
		})
	}
}
