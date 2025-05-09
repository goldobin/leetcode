package find_median_from_data_stream

import (
	"fmt"
	"math/rand"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MedianFinder(t *testing.T) {
	tests := []struct {
		name string
		ops  []op
	}{
		{
			name: "case 0.1 empty",
			ops:  []op{findMedianOp(0)},
		},
		{
			name: "case 0.2 one entry",
			ops: []op{
				addNumOp(1),
				findMedianOp(1),
			},
		},
		{
			name: "case 0.2 one entry",
			ops: []op{
				addNumOp(2),
				findMedianOp(2),
			},
		},
		{
			name: "case 0.3 two entries",
			ops: []op{
				addNumOp(2),
				addNumOp(3),
				findMedianOp(2.5),
			},
		},
		{
			name: "case 0.4 three entries",
			ops: []op{
				addNumOp(2),
				addNumOp(3),
				addNumOp(10),
				findMedianOp(3),
			},
		},
		{
			name: "case 1.1 leetcode",
			ops: []op{
				addNumOp(1),
				addNumOp(2),
				findMedianOp(1.5),
				addNumOp(3),
				findMedianOp(2),
			},
		},
		{
			name: "case 2.1.1",
			ops:  numsToOps([]int{-11000000, 1, 101, 900, 90000}),
		},
		{
			name: "case 2.1.2 reverse",
			ops:  numsToOps([]int{90000, 900, 101, 1, -11000000}),
		},
		{
			name: "case 2.1.3 shuffle",
			ops:  numsToOps([]int{90000, 101, -11000000, 1, 900}),
		},
		{
			name: "case 2.1.4 duplicates",
			ops:  numsToOps([]int{90000, 101, 900, 101, -11000000, 1, 900}),
		},
		{
			name: "case 2.1.5 duplicates",
			ops:  numsToOps([]int{-11000000, 90000, 900, 101, -11000000, 1, 900}),
		},
		{
			name: "case 2.1.6 duplicates",
			ops:  numsToOps([]int{90000, 900, 101, -11000000, 1, 900, 90000}),
		},
		{
			name: "case 2.2.1",
			ops:  numsToOps([]int{-11000000, 1, 101, 900, 901, 90000}),
		},
		{
			name: "case 2.2.2 reverse",
			ops:  numsToOps([]int{90000, 901, 900, 101, 1, -11000000}),
		},
		{
			name: "case 2.2.3 shuffle",
			ops:  numsToOps([]int{90000, -11000000, 101, 900, 1, 901}),
		},
		{
			name: "case 4.1 random positive",
			ops: func() []op {
				nums := make([]int, 100)
				generateRandNums(nums, 0, 100)
				return numsToOps(nums)
			}(),
		},
		{
			name: "case 4.2 random negative",
			ops: func() []op {
				nums := make([]int, 100)
				generateRandNums(nums, -100, 0)
				return numsToOps(nums)
			}(),
		},
		{
			name: "case 4.3 random 100",
			ops: func() []op {
				nums := make([]int, 100)
				generateRandNums(nums, -100, 100)
				return numsToOps(nums)
			}(),
		},
		{
			name: "case 4.4 random 200",
			ops: func() []op {
				nums := make([]int, 200)
				generateRandNums(nums, -100, 100)
				return numsToOps(nums)
			}(),
		},
		{
			name: "case 4.5 random 1000",
			ops: func() []op {
				nums := make([]int, 1000)
				generateRandNums(nums, -100, 100)
				return numsToOps(nums)
			}(),
		},
		{
			name: "case 4.6 random 49999",
			ops: func() []op {
				nums := make([]int, 49999)
				generateRandNums(nums, -100000, 100000)
				return numsToOps(nums)
			}(),
		},
		{
			name: "case 4.7 random 50000",
			ops: func() []op {
				nums := make([]int, 50000)
				generateRandNums(nums, -100000, 100000)
				return numsToOps(nums)
			}(),
		},
		{
			name: "case 5.1 ordered 25`",
			ops: func() []op {
				nums := []int{
					-963, -856, -785, -655, -595, -590, -558, -273, -256, -133, -76,
					-25, 44, 84, 103, 328, 361, 475, 628, 646, 684, 695, 815, 901, 990,
				}
				// slices.Sort(nums)
				fmt.Println(nums)
				return numsToOps2(nums)
			}(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mf := Constructor()
			for _, op := range test.ops {
				op(t, &mf)
			}
		})
	}
}

func Benchmark_MedianFinder(b *testing.B) {
	b.Run("odd", func(b *testing.B) {
		nums := make([]int, 4999)
		generateRandNums(nums, -100000, 100000)
		mf := Constructor()
		var m float64

		for b.Loop() {
			for _, num := range nums {
				mf.AddNum(num)
			}
			m = mf.FindMedian()
		}
		b.Logf("median: %v", m)
	})

	b.Run("even", func(b *testing.B) {
		nums := make([]int, 5000)
		generateRandNums(nums, -100000, 100000)
		mf := Constructor()
		var m float64

		for b.Loop() {
			for _, num := range nums {
				mf.AddNum(num)
			}
			m = mf.FindMedian()
		}
		b.Logf("median: %v", m)
	})
}

func median(nums []int) float64 {
	if len(nums) == 0 {
		return 0
	}

	slices.Sort(nums)
	if len(nums)%2 == 0 {
		return float64(nums[len(nums)/2-1]+nums[len(nums)/2]) / 2.0
	}

	return float64(nums[len(nums)/2])
}

func Test_median(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want float64
	}{
		{
			name: "case 0.1 empty",
			nums: []int{},
			want: 0,
		},
		{
			name: "case 0.2 one entry",
			nums: []int{1},
			want: 1,
		},
		{
			name: "case 0.3 one entry",
			nums: []int{2},
			want: 2,
		},
		{
			name: "case 0.4 one entry",
			nums: []int{-100},
			want: -100,
		},
		{
			name: "case 1.1 two entries",
			nums: []int{3, 4},
			want: 3.5,
		},
		{
			name: "case 1.2 two entries",
			nums: []int{3, 5},
			want: 4,
		},
		{
			name: "case 1.3 two entries",
			nums: []int{-100, 100},
			want: 0,
		},
		{
			name: "case 2.1 three entries",
			nums: []int{-100, 100, 101},
			want: 100,
		},
		{
			name: "case 2.2 three entries",
			nums: []int{99, 100, 101},
			want: 100,
		},
		{
			name: "case 2.3 three entries",
			nums: []int{-11000000, 1, 101},
			want: 1,
		},
		{
			name: "case 3.1 four entries",
			nums: []int{-11000000, 1, 101, 900},
			want: 51,
		},
		{
			name: "case 4.1 more entries",
			nums: []int{-11000000, 1, 101, 900, 90000},
			want: 101,
		},
		{
			name: "case 4.1 more entries",
			nums: []int{-11000000, 1, 101, 900, 90000, 901},
			want: 500.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := median(slices.Clone(tt.nums))
			assert.InDelta(t, tt.want, got, 0.0001)
		})
	}
}

type op func(t *testing.T, mf *MedianFinder)

func addNumOp(num int) op {
	return func(t *testing.T, mf *MedianFinder) {
		mf.AddNum(num)
	}
}

func findMedianOp(want float64) op {
	return func(t *testing.T, mf *MedianFinder) {
		got := mf.FindMedian()
		assert.InDelta(t, want, got, 0.001)
	}
}

func numsToAddNumOps(dst []op, nums []int) int {
	n := min(len(dst), len(nums))
	for i := 0; i < n; i++ {
		dst[i] = addNumOp(nums[i])
	}

	return n
}

func generateRandNums(dst []int, min, max int) {
	n := max - min
	if n < 1 {
		panic("min >= max")
	}

	for i := 0; i < len(dst); i++ {
		dst[i] = rand.Intn(n) + min
	}
}

func numsToOps(nums []int) []op {
	n := len(nums)
	ops := make([]op, n+1)
	numsToAddNumOps(ops, nums)
	ops[n] = findMedianOp(median(nums))
	return ops
}

func numsToOps2(nums []int) []op {
	n := len(nums)
	ops := make([]op, n*2)
	for i := 0; i < n; i++ {
		ops[i*2] = addNumOp(nums[i])
		ops[i*2+1] = findMedianOp(median(nums[:i+1]))
	}

	return ops
}
