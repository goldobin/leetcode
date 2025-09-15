package summary_ranges

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_summaryRanges(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []string
	}{
		{
			name: "case 0.1 empty nums",
		},
		{
			name: "case 0.2.1 1 element",
			nums: []int{0},
			want: []string{"0"},
		},
		{
			name: "case 0.2.2 1 element",
			nums: []int{10},
			want: []string{"10"},
		},
		{
			name: "case 0.2.3 distinct elements",
			nums: []int{10, 12, 14, 16},
			want: []string{"10", "12", "14", "16"},
		},
		{
			name: "case 1.1 leetcode",
			nums: []int{0, 1, 2, 4, 5, 7},
			want: []string{"0->2", "4->5", "7"},
		},
		{
			name: "case 1.2 leetcode",
			nums: []int{0, 2, 3, 4, 6, 8, 9},
			want: []string{"0", "2->4", "6", "8->9"},
		},
		{
			name: "case 1.3 negatives",
			nums: []int{-9, -8, -6, -4, -3, -2, 0},
			want: []string{"-9->-8", "-6", "-4->-2", "0"},
		},
		{
			name: "case 1.4 negatives and positives",
			nums: []int{-9, -8, -6, -4, -3, -2, 0, 1, 2, 3, 4, 10},
			want: []string{"-9->-8", "-6", "-4->-2", "0->4", "10"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := summaryRanges(tt.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}
