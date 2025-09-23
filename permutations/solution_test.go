package permutations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_permute(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "case 0.1 empty nums",
		},
		{
			name: "case 0.2 one num",
			nums: []int{1},
			want: [][]int{{1}},
		},
		{
			name: "case 0.3 two nums",
			nums: []int{1, 2},
			want: [][]int{{1, 2}, {2, 1}},
		},
		{
			name: "case 0.4 three nums",
			nums: []int{1, 2, 3},
			want: [][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := permute(tt.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Benchmark_permute(b *testing.B) {
	b.ReportAllocs()
	nums := make([]int, 6)
	for i := 0; i < len(nums); i++ {
		nums[i] = i + 1
	}

	for b.Loop() {
		permute(nums)
	}
}
