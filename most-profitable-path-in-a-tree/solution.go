package most_profitable_path_in_a_tree

import (
	"math"
)

type node struct {
	parent       int
	children     map[int]struct{}
	bobVisited   bool
	bobVisitStep int
}

func mostProfitablePath(edges [][]int, bobStartNode int, rewards []int) int {
	var (
		nodeCount = len(rewards)
		tree      = make([]node, nodeCount)
	)

	buildTree(tree, edges)
	markBobPath(tree, bobStartNode)

	s := solver{
		rewards,
		tree,
	}

	result := s.solve(0)
	return result
}

func buildTree(resultTree []node, edges [][]int) {
	for _, edge := range edges {
		var (
			from = edge[0]
			to   = edge[1]
		)

		if resultTree[from].children == nil {
			resultTree[from].children = make(map[int]struct{})
		}
		resultTree[from].children[to] = struct{}{}

		if resultTree[to].children == nil {
			resultTree[to].children = make(map[int]struct{})
		}
		resultTree[to].children[from] = struct{}{}
	}

	var (
		visits = make([]bool, len(resultTree))
		queue  = make([]int, 0)
	)
	visits[0] = true
	resultTree[0].parent = -1
	queue = append(queue, 0)

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		children := resultTree[p].children
		for c := range children {
			if visits[c] {
				delete(children, c)
				continue
			}
			visits[c] = true
			resultTree[c].parent = p
			queue = append(queue, c)
		}

		if len(children) == 0 {
			resultTree[p].children = nil
		}
	}
}

func markBobPath(tree []node, startNode int) {
	var (
		step = 0
		next = startNode
	)

	for {
		tree[next].bobVisited = true
		tree[next].bobVisitStep = step
		step += 1
		if next == 0 {
			break
		}
		next = tree[next].parent
		if next < 0 {
			break
		}
	}
}

type solver struct {
	rewards []int
	tree    []node
}

func (s solver) solve(aliceStartNode int) int {
	return s.iterate(aliceStartNode, 0)
}

func (s solver) iterate(i int, step int) int {
	var (
		n            = s.tree[i]
		reward       = s.rewards[i]
		actualReward int
	)

	if !n.bobVisited {
		actualReward = reward
	} else if step == n.bobVisitStep {
		actualReward = reward / 2
	} else if step < n.bobVisitStep {
		actualReward = reward
	} else {
		actualReward = 0
	}

	if len(n.children) == 0 {
		return actualReward
	}

	maxSubReward := math.MinInt32
	for c := range n.children {
		subReward := s.iterate(c, step+1)
		if subReward > maxSubReward {
			maxSubReward = subReward
		}
	}

	return actualReward + maxSubReward
}
