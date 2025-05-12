package max_stack

import (
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

func popMaxOp(want int) op {
	return func(t *testing.T, s *MaxStack) {
		v := s.PopMax()
		assert.Equal(t, want, v)
	}
}
