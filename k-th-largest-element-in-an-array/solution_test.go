package k_th_largest_element_in_an_array

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findKthLargest(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "case 0.1",
			nums: []int{3},
			k:    1,
			want: 3,
		},

		{
			name: "case 1.1 leetcode",
			nums: []int{3, 2, 1, 5, 6, 4},
			k:    2,
			want: 5,
		},
		{
			name: "case 1.2 leetcode",
			nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:    4,
			want: 4,
		},
		{
			name: "case 1.3",
			nums: []int{3, 2, 1, 5, 6, 4},
			k:    6,
			want: 1,
		},
		{
			name: "case 1.4",
			nums: []int{5, 5, 5, 5, 6, 4},
			k:    6,
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findKthLargest(tt.nums, tt.k)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_peek_on_empty_heap(t *testing.T) {
	var h heap

	_, ok := h.peek()
	assert.False(t, ok)
}

func Test_pull_on_empty_heap(t *testing.T) {
	var h heap

	_, ok := h.pull()
	assert.False(t, ok)
}

func Test_heap_push(t *testing.T) {
	tests := []struct {
		name string
		nums []int
	}{
		{
			name: "case 0.1",
			nums: []int{1},
		},

		{
			name: "case 0.2",
			nums: []int{1, 2},
		},
		{
			name: "case 0.3",
			nums: []int{1, 2, 3},
		},
		{
			name: "case 0.4",
			nums: []int{1, 2, 3, 3},
		},
		{
			name: "case 0.5",
			nums: []int{2, 1, 4, 3},
		},
		{
			name: "case 0.6",
			nums: []int{2, 1},
		},
		{
			name: "case 0.7",
			nums: []int{3, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var h heap

			for i, n := range tt.nums {
				wantPeek := maxEl(tt.nums[:i+1])

				h.push(n)
				got, gotOK := h.peek()
				if assert.True(t, gotOK) {
					assert.Equal(t, wantPeek, got)
				}
			}

			sortedNums := cloneAndSort(tt.nums)
			for _, n := range sortedNums {
				got, gotOK := h.pull()
				if assert.True(t, gotOK) {
					assert.Equal(t, n, got)
				}
			}
		})
	}
}

func cloneAndSort(ns []int) []int {
	sns := slices.Clone(ns)
	slices.Sort(sns)
	slices.Reverse(sns)
	return sns
}

func maxEl(ns []int) int {
	if len(ns) == 0 {
		panic("slice is empty")
	}

	sns := cloneAndSort(ns)
	return sns[0]
}
