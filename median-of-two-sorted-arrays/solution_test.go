package median_of_two_sorted_arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMedianSortedArrays(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  float64
	}{
		{
			name:  "case 0.1.1",
			nums1: nil,
			nums2: []int{2},
			want:  2.0,
		},
		{
			name:  "case 0.1.2",
			nums1: nil,
			nums2: []int{2, 2},
			want:  2.0,
		},
		{
			name:  "case 0.1.3",
			nums1: nil,
			nums2: []int{2, 3},
			want:  2.5,
		},
		{
			name:  "case 0.1.4",
			nums1: nil,
			nums2: []int{2, 3, 5},
			want:  3,
		},
		{
			name:  "case 0.1.5",
			nums1: nil,
			nums2: []int{2, 3, 4, 5},
			want:  3.5,
		},
		{
			name:  "case 0.2.1",
			nums1: []int{3},
			nums2: nil,
			want:  3.0,
		},
		{
			name:  "case 0.2.2",
			nums1: []int{3, 3},
			nums2: nil,
			want:  3.0,
		},
		{
			name:  "case 0.2.3",
			nums1: []int{3, 4},
			nums2: nil,
			want:  3.5,
		},
		{
			name:  "case 0.2.4",
			nums1: []int{3, 4, 5},
			nums2: nil,
			want:  4,
		},
		{
			name:  "case 0.2.5",
			nums1: []int{3, 4, 5, 6},
			nums2: nil,
			want:  4.5,
		},
		{
			name:  "case 1.1 odd leetcode",
			nums1: []int{1, 3},
			nums2: []int{2},
			want:  2,
		},
		{
			name:  "case 1.2 odd",
			nums1: []int{1, 2, 4, 5},
			nums2: []int{3},
			want:  3,
		},
		{
			name:  "case 1.3 odd",
			nums1: []int{1, 3, 4, 5},
			nums2: []int{2},
			want:  3,
		},
		{
			name:  "case 2.2 leetcode",
			nums1: []int{1, 2},
			nums2: []int{3, 4},
			want:  2.5,
		},
		{
			name:  "case 2.3",
			nums1: []int{1, 3, 4, 6, 7}, // 1, 2, 3, [4, 5], 6, 7, 8
			nums2: []int{2, 5, 8},
			want:  4.5,
		},
		{
			name:  "case 2.4",
			nums1: []int{0, 1, 3, 4, 6, 7, 17, 18, 19, 20}, // 0, 1, 2, 3, 4, 5, [6], 7, 8, 17, 18, 19, 20
			nums2: []int{2, 5, 8},
			want:  6,
		},
		{
			name:  "case 2.5 leetcode",
			nums1: []int{1, 2, 3, 4, 5},
			nums2: []int{6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17},
			want:  9,
		},
		{
			name:  "case 2.6 leetcode",
			nums1: []int{2},
			nums2: []int{1, 3, 4, 5, 6},
			want:  3.5,
		},
		{
			name:  "case 2.7 leetcode",
			nums1: []int{1, 2, 6, 7}, // 1, 2, 3, [4, 5], 6, 7, 8
			nums2: []int{3, 4, 5, 8},
			want:  4.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMedianSortedArrays(tt.nums1, tt.nums2)
			assert.InEpsilonf(t, tt.want, got, 0.001, "expected %v, got %v", tt.want, got)
		})
	}
}

func Test_findMedianSortedPanicsIdBothArraysNilOrEmpty(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
	}{
		{
			nums1: nil,
			nums2: nil,
		},
		{
			nums1: []int{},
			nums2: nil,
		},
		{
			nums1: nil,
			nums2: []int{},
		},
		{
			nums1: []int{},
			nums2: []int{},
		},
	}

	for _, tt := range tests {
		assert.Panics(t, func() {
			findMedianSortedArrays(tt.nums1, tt.nums2)
		})
	}
}

func Test_median(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want float64
	}{
		{
			name: "case 1.1",
			nums: []int{2},
			want: 2.0,
		},
		{
			name: "case 1.2",
			nums: []int{2, 3},
			want: 2.5,
		},
		{
			name: "case 1.3",
			nums: []int{2, 3, 4},
			want: 3.0,
		},
		{
			name: "case 1.4",
			nums: []int{2, 3, 3, 4},
			want: 3.0,
		},
		{
			name: "case 1.5",
			nums: []int{4, 5, 6, 7, 8, 10, 12},
			want: 7,
		},
		{
			name: "case 1.6",
			nums: []int{4, 5, 6, 7, 9, 9, 10, 12},
			want: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := median(tt.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}
