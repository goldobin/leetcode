package find_first_and_last_position_of_element_in_sorted_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_searchRange(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "case 0.1",
			nums:   []int{},
			target: 1,
			want:   []int{-1, -1},
		},
		{
			name:   "case 0.2",
			nums:   []int{2},
			target: 2,
			want:   []int{0, 0},
		},
		{
			name:   "case 0.3",
			nums:   []int{2, 3},
			target: 3,
			want:   []int{1, 1},
		},
		{
			name:   "case 0.4",
			nums:   []int{2, 3},
			target: 2,
			want:   []int{0, 0},
		},
		{
			name:   "case 0.5",
			nums:   []int{2, 3, 5},
			target: 3,
			want:   []int{1, 1},
		},
		{
			name:   "case 0.6",
			nums:   []int{2, 3, 3, 5},
			target: 3,
			want:   []int{1, 2},
		},
		{
			name:   "case 1.1 leetcode",
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 8,
			want:   []int{3, 4},
		},
		{
			name:   "case 1.2 leetcode",
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 6,
			want:   []int{-1, -1},
		},
		{
			name:   "case 1.3 leetcode",
			nums:   []int{5, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 8, 8, 10},
			target: 7,
			want:   []int{1, 34},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchRange(tt.nums, tt.target)
			if assert.Len(t, got, 2) {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func Test_binarySearchLR(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		wantL  int
		wantR  int
	}{
		{
			name:   "case 0.1",
			nums:   []int{},
			target: 1,
			wantL:  -1,
			wantR:  -1,
		},
		{
			name:   "case 0.2",
			nums:   []int{2},
			target: 1,
			wantL:  -1,
			wantR:  -1,
		},
		{
			name:   "case 0.3",
			nums:   []int{2, 3},
			target: 2,
			wantL:  0,
			wantR:  0,
		},
		{
			name:   "case 0.4",
			nums:   []int{2, 3},
			target: 3,
			wantL:  1,
			wantR:  1,
		},
		{
			name:   "case 0.5",
			nums:   []int{2, 3, 4},
			target: 3,
			wantL:  1,
			wantR:  1,
		},
		{
			name:   "case 0.6",
			nums:   []int{2, 3, 4, 4, 5},
			target: 4,
			wantL:  2,
			wantR:  3,
		},
		{
			name:   "case 0.6",
			nums:   []int{2, 3, 4, 4, 4, 4, 4, 4, 4, 5},
			target: 4,
			wantL:  2,
			wantR:  8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotL := binarySearchL(tt.nums, 0, len(tt.nums), tt.target)
			gotR := binarySearchR(tt.nums, 0, len(tt.nums), tt.target)
			assert.Equal(t, tt.wantL, gotL)
			assert.Equal(t, tt.wantR, gotR)
		})
	}
}
