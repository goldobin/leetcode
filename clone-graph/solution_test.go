package clone_graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_cloneGraph(t *testing.T) {
	tests := []struct {
		name string
		node *Node
	}{
		{
			name: "case 0.1",
			node: &Node{
				Val: 1,
			},
		},
		{
			name: "case 1.1 leetcode",
			node: func() *Node {
				node1 := &Node{Val: 1}
				node2 := &Node{Val: 2}
				node3 := &Node{Val: 3}
				node4 := &Node{Val: 4}

				node1.Neighbors = []*Node{node2, node4}
				node2.Neighbors = []*Node{node1, node3}
				node3.Neighbors = []*Node{node4, node2}
				node4.Neighbors = []*Node{node1, node3}

				return node1
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cloneGraph(tt.node)
			if !assert.NotNil(t, got) || !assert.False(t, tt.node == got) {
				return
			}
			assert.Equal(t, tt.node, got)
		})
	}
}
