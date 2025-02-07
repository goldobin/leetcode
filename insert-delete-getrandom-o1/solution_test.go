package insert_delete_getrandom_o1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RandomizedSet(t *testing.T) {
	tests := []struct {
		name string
		ops  []operation
	}{
		{
			name: "case 0.1",
			ops: []operation{
				insert(1, true),
				checkRandom([]int{1}, 5),
			},
		},
		{
			name: "case 0.2",
			ops: []operation{
				remove(1, false),
				insert(1, true),
				remove(1, true),
				insert(2, true),
				checkRandom([]int{2}, 10),
			},
		},
		{
			name: "case 1.1 leetcode",
			ops: []operation{
				insert(1, true),
				remove(2, false),
				insert(2, true),
				checkRandom([]int{1, 2}, 10),
				remove(1, true),
				insert(2, false),
				checkRandom([]int{2}, 10),
			},
		},
		{
			name: "case 1.2",
			ops: []operation{
				insert(0, true),
				insert(1, true),
				insert(2, true),
				insert(3, true),
				insert(4, true),
				insert(5, true),
				insert(6, true),
				insert(7, true),
				remove(0, true),
				remove(1, true),
				remove(2, true),
				remove(3, true),
				remove(4, true),
				remove(5, true),
				remove(6, true),
				remove(7, true),
				remove(0, false),
				remove(1, false),
				remove(2, false),
				remove(3, false),
				remove(4, false),
				remove(5, false),
				remove(6, false),
				remove(7, false),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Constructor()
			for _, op := range tt.ops {
				if !op(t, &s) {
					return
				}
			}
		})
	}
}

type operation func(t *testing.T, set *RandomizedSet) bool

func insert(val int, want bool) operation {
	return func(t *testing.T, s *RandomizedSet) bool {
		got := s.Insert(val)
		return assert.Equal(t, want, got)
	}
}

func remove(val int, want bool) operation {
	return func(t *testing.T, set *RandomizedSet) bool {
		got := set.Remove(val)
		return assert.Equal(t, want, got)
	}
}

func checkRandom(wantOneOf []int, trials int) operation {
	return func(t *testing.T, s *RandomizedSet) bool {
		for i := 0; i < trials; i++ {
			got := s.GetRandom()
			if !assert.Contains(t, wantOneOf, got) {
				return false
			}
		}

		return true
	}
}
