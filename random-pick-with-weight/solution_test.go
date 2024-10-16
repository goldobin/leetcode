package random_pick_with_weight

import (
	"cmp"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_buildTree(t *testing.T) {
	tests := []struct {
		name string
		es   []int
		want *node[int]
	}{
		{
			name: "empty array",
			es:   []int{},
			want: nil,
		},
		{
			name: "single element",
			es:   []int{1},
			want: &node[int]{value: 1},
		},
		{
			name: "two elements",
			es:   []int{1, 2},
			want: &node[int]{
				value: 2,
				left: &node[int]{
					value: 1,
				},
			},
		},
		{
			name: "three elements",
			es:   []int{1, 2, 3},
			want: &node[int]{
				value: 2,
				left: &node[int]{
					value: 1,
				},
				right: &node[int]{
					value: 3,
				},
			},
		},
		{
			name: "four elements",
			es:   []int{1, 2, 3, 4},
			want: &node[int]{
				value: 3,
				left: &node[int]{
					value: 2,
					left: &node[int]{
						value: 1,
					},
				},
				right: &node[int]{
					value: 4,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildTree(tt.es)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_node_ceilFunc(t *testing.T) {
	tests := []struct {
		name string
		node *node[int]
		x    int
		want *node[int]
	}{
		{
			name: "single node is ceil",
			node: &node[int]{
				value: 0,
			},
			x: 1,
			want: &node[int]{
				value: 0,
			},
		},
		{
			name: "single node is not a ceil",
			node: &node[int]{
				value: 1,
			},
			x:    0,
			want: nil,
		},
		{
			name: "two nodes: root and left, no ceil",
			node: &node[int]{
				value: 3,
				left: &node[int]{
					value: 2,
				},
			},
			x:    1,
			want: nil,
		},
		{
			name: "two nodes: root and left, left is ceil",
			node: &node[int]{
				value: 2,
				left: &node[int]{
					value: 0,
				},
			},
			x: 1,
			want: &node[int]{
				value: 0,
			},
		},
		{
			name: "two nodes: root and right, no ceil",
			node: &node[int]{
				value: 1,
				right: &node[int]{
					value: 2,
				},
			},
			x:    0,
			want: nil,
		},
		{
			name: "two nodes: root and right, root is ceil",
			node: &node[int]{
				value: 1,
				right: &node[int]{
					value: 3,
				},
			},
			x: 2,
			want: &node[int]{
				value: 1,
				right: &node[int]{
					value: 3,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.node.ceilFunc(tt.x, cmp.Compare[int])
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_Solution_PickIndex(t *testing.T) {
	var (
		eps     = 0.05
		samples = 1000
		tests   = [][]int{
			{1, 3, 4},
			{1, 3, 4, 5},
			{1, 20},
			{1, 1, 1, 1},
			{2, 2, 2, 2},
			{1, 2, 4, 4, 2, 1},
			{1, 2, 5, 7, 5, 3, 1},
		}
	)

	for _, ws := range tests {
		var (
			name     = fmt.Sprintf("weights=%v", ws)
			sum      = 0
			wantDist = make(map[int]float64)
		)

		for _, v := range ws {
			sum += v
		}

		for i, v := range ws {
			wantDist[i] = float64(v) / float64(sum)
		}

		t.Run(name, func(t *testing.T) {
			var (
				s         = Constructor(ws)
				gotCounts = make(map[int]int)
				gotDist   = make(map[int]float64)
			)

			for i := 0; i < samples; i++ {
				gotIx := s.PickIndex()
				assert.Greater(t, gotIx, -1)
				assert.Less(t, gotIx, len(ws))

				gotCounts[gotIx]++
			}

			gotSum := 0
			for _, v := range gotCounts {
				gotSum += v
			}

			for k, v := range gotCounts {
				gotDist[k] = float64(v) / float64(gotSum)
			}

			for k, v := range wantDist {
				assert.InDelta(t, v, gotDist[k], eps)
			}
		})
	}
}
