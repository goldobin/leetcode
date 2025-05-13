package max_stack

import (
	"math"
	"math/rand/v2"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MaxStack_PanicOnEmptyStack(t *testing.T) {
	tests := []struct {
		name string
		ops  []op
	}{
		{
			name: "case 0.1",
			ops: []op{
				popOp(0),
			},
		},
		{
			name: "case 0.2",
			ops: []op{
				topOp(0),
			},
		},
		{
			name: "case 0.3",
			ops: []op{
				peekMaxOp(0),
			},
		},
		{
			name: "case 0.4",
			ops: []op{
				popMaxOp(0),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Constructor()
			assert.Panics(t, func() {
				for _, op := range tt.ops {
					op(t, &s)
				}
			})
		})
	}
}

func Test_MaxStack(t *testing.T) {
	tests := []struct {
		name string
		ops  []op
	}{
		{
			name: "case 1.1 leetcode",
			ops: []op{
				pushOp(5, 1, 5),
				topOp(5),
				popMaxOp(5),
				topOp(1),
				peekMaxOp(5),
				popOp(1),
				topOp(5),
			},
		},
		{
			name: "case 2.1 sequential asc desc",
			ops: []op{
				pushOp(1, 2, 3, 3, 4, 5, 4, 3, 2, 1, 1, 5),
				peekMaxOp(5),
				popOp(5),
				peekMaxOp(5),
				topOp(1),
				popMaxOp(5),
				popOp(1, 1, 2, 3, 4),
				topOp(4),
				popMaxOp(4),
				topOp(3),
				popMaxOp(3),
				topOp(3),
				peekMaxOp(3),
				popOp(3, 2, 1),
			},
		},
		{
			name: "case 2.2 push pop all",
			ops: []op{
				pushOp(1, 2, 3, 4, 5),
				popOp(5, 4, 3, 2, 1),
				lenOp(0),
			},
		},
		{
			name: "case 2.3 push pop max all",
			ops: []op{
				pushOp(1, 2, 3, 4, 5),
				popMaxOp(5, 4, 3, 2, 1),
				lenOp(0),
			},
		},
		{
			name: "case 2.4 push pop the same",
			ops: []op{
				pushOp(10, 10, 10, 10, 10),
				popOp(10, 10, 10, 10, 10),
				lenOp(0),
			},
		},
		{
			name: "case 2.5 push pop max the same",
			ops: []op{
				pushOp(11, 11, 11, 11, 11),
				popMaxOp(11, 11, 11, 11, 11),
				lenOp(0),
			},
		},
		{
			name: "case 2.6 push pop max the same mixed",
			ops: []op{
				pushOp(12, 12, 12, 12, 12, 12),
				popOp(12, 12, 12),
				popMaxOp(12, 12, 12),
				lenOp(0),
			},
		},
		{
			name: "case 2.7 push pop max the same mixed",
			ops: []op{
				pushOp(12, 12, 12, 12, 12, 12),
				popMaxOp(12, 12, 12),
				popOp(12, 12, 12),
				lenOp(0),
			},
		},
		{
			name: "case 3.1 push pop max random 100 batch 5",
			ops: []op{
				pushPopRandom(100, 5, -intPow10(7), intPow10(7)),
				lenOp(0),
			},
		},
		{
			name: "case 3.2 push pop max random 1000 batch 50",
			ops: []op{
				pushPopRandom(1000, 50, -intPow10(7), intPow10(7)),
				lenOp(0),
			},
		},
		{
			name: "case 3.3 push pop max random 1000 batch 500",
			ops: []op{
				pushPopRandom(1000, 100, -intPow10(7), intPow10(7)),
				lenOp(0),
			},
		},
		{
			name: "case 3.4 push pop max random 10000 batch 50",
			ops: []op{
				pushPopRandom(10000, 10, -intPow10(7), intPow10(7)),
				lenOp(0),
			},
		},
		{
			name: "case 3.5 push pop max random 10000 batch 50",
			ops: []op{
				pushPopRandom(100000, 1, -intPow10(7), intPow10(7)),
				lenOp(0),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Constructor()
			for _, op := range tt.ops {
				op(t, &s)
			}
		})
	}
}

type op func(t *testing.T, s *MaxStack)

func pushOp(vs ...int) op {
	return func(_ *testing.T, s *MaxStack) {
		for _, v := range vs {
			s.Push(v)
		}
	}
}

func popOp(wants ...int) op {
	return func(t *testing.T, s *MaxStack) {
		for _, want := range wants {
			v := s.Pop()
			assert.Equal(t, want, v)
		}
	}
}

func topOp(want int) op {
	return func(t *testing.T, s *MaxStack) {
		v := s.Top()
		assert.Equal(t, want, v)
	}
}

func peekMaxOp(want int) op {
	return func(t *testing.T, s *MaxStack) {
		v := s.PeekMax()
		assert.Equal(t, want, v)
	}
}

func popMaxOp(wants ...int) op {
	return func(t *testing.T, s *MaxStack) {
		for _, want := range wants {
			v := s.PopMax()
			assert.Equal(t, want, v)
		}
	}
}

func lenOp(want int) op {
	return func(t *testing.T, s *MaxStack) {
		assert.Equal(t, want, s.Len())
	}
}

func pushPopRandom(n int, batchLen int, min int, max int) op {
	return func(t *testing.T, s *MaxStack) {
		rMax := max - min
		nums := make([]int, batchLen)

		for i := 0; i < n; i++ {
			maxNum := math.MinInt
			maxNumIndex := -1
			for j := 0; j < batchLen; j++ {
				nums[j] = rand.IntN(rMax) + min
				if maxNum < nums[j] {
					maxNum = nums[j]
					maxNumIndex = j
				}
			}

			for j := 0; j < batchLen; j++ {
				s.Push(nums[j])
			}

			assert.Equal(t, maxNum, s.PeekMax())
			assert.Equal(t, maxNum, s.PopMax())

			for j := batchLen - 1; j >= 0; j-- {
				if j == maxNumIndex {
					continue
				}
				v := s.Pop()
				assert.Equal(t, nums[j], v)
			}
		}
	}
}

func intPow10(p int) int {
	result := 1
	for i := 0; i < p; i++ {
		result *= 10
	}
	return result
}
