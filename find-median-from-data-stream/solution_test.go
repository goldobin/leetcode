package find_median_from_data_stream

import (
	"fmt"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PeekOnEmptyHeap(t *testing.T) {
	var h heap

	_, ok := h.peek()
	assert.False(t, ok)
}

func Test_PullOnEmptyHeap(t *testing.T) {
	var h heap

	_, ok := h.pull()
	assert.False(t, ok)
}

func Test_HeapPush(t *testing.T) {
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

		{
			name: "case 0.8",
			nums: []int{3, 3, 2, 1, 1},
		},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("min heap %s", tt.name)
		t.Run(name, func(t *testing.T) {
			var h heap = newMinHeap()

			for i, n := range tt.nums {
				wantPeek := minElement(tt.nums[:i+1])

				h.push(n)
				got, gotOK := h.peek()
				if assert.True(t, gotOK) {
					assert.Equal(t, wantPeek, got)
				}
			}

			sortedNums := cloneAndSortAsc(tt.nums)
			for _, n := range sortedNums {
				got, gotOK := h.pull()
				if assert.True(t, gotOK) {
					assert.Equal(t, n, got)
				}
			}
		})
	}

	for _, tt := range tests {
		name := fmt.Sprintf("max heap %s", tt.name)
		t.Run(name, func(t *testing.T) {
			var h heap = newMaxHeap()

			for i, n := range tt.nums {
				wantPeek := maxElement(tt.nums[:i+1])

				h.push(n)
				got, gotOK := h.peek()
				if assert.True(t, gotOK) {
					assert.Equal(t, wantPeek, got)
				}
			}

			sortedNums := cloneAndSortDesc(tt.nums)
			for _, n := range sortedNums {
				got, gotOK := h.pull()
				if assert.True(t, gotOK) {
					assert.Equal(t, n, got)
				}
			}
		})
	}
}

func cloneAndSortAsc(ns []int) []int {
	result := slices.Clone(ns)
	slices.Sort(result)
	return result
}

func cloneAndSortDesc(ns []int) []int {
	result := slices.Clone(ns)
	slices.Sort(result)
	slices.Reverse(result)
	return result
}

func minElement(ns []int) int {
	if len(ns) == 0 {
		panic("slice is empty")
	}

	result := cloneAndSortAsc(ns)
	return result[0]
}

func maxElement(ns []int) int {
	if len(ns) == 0 {
		panic("slice is empty")
	}

	sns := cloneAndSortDesc(ns)
	return sns[0]
}
