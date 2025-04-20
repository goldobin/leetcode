package most_profitable_path_in_a_tree

import (
	"maps"
	"math"
	"slices"
)

func mostProfitablePath(edges [][]int, bobStartNode int, rewards []int) int {
	var (
		nodeCount = len(rewards)
		children  = make([][]int, nodeCount)
		parents   = make([]int, nodeCount)
		bobVisits = make([]int, nodeCount)
	)

	buildChildren(children, edges)
	buildParents(parents, children)
	buildBobVisits(bobVisits, parents, bobStartNode)

	s := solver{
		rewards,
		children,
		bobVisits,
	}

	return s.solve(0)
}

func buildChildren(result [][]int, edges [][]int) {
	var (
		nodeCount     = len(result)
		adjacentNodes = make([]map[int]struct{}, nodeCount)
		visits        = make([]bool, nodeCount)
		queue         = make([]int, 0, len(edges))
	)

	for _, edge := range edges {
		var (
			from = edge[0]
			to   = edge[1]
		)

		if adjacentNodes[from] == nil {
			adjacentNodes[from] = make(map[int]struct{}, nodeCount)
		}
		adjacentNodes[from][to] = struct{}{}

		if adjacentNodes[to] == nil {
			adjacentNodes[to] = make(map[int]struct{}, nodeCount)
		}
		adjacentNodes[to][from] = struct{}{}
	}

	queue = append(queue, 0)
	visits[0] = true

	for len(queue) > 0 {
		currentParent := queue[0]
		queue = queue[1:]

		adjNodes := adjacentNodes[currentParent]
		for v := range adjNodes {
			if !visits[v] {
				continue
			}
			delete(adjNodes, v)
		}

		var childNodes []int = slices.Collect(maps.Keys(adjNodes))
		result[currentParent] = childNodes
		for _, n := range childNodes {
			visits[n] = true
		}
		queue = append(queue, childNodes...)
	}
}

func buildParents(result []int, children [][]int) {
	for i := 0; i < len(result); i++ {
		result[i] = -1
	}

	for i := 0; i < len(children); i++ {
		for _, v := range children[i] {
			result[v] = i
		}
	}
}

func buildBobVisits(visits []int, parents []int, startNode int) {
	var (
		step = 0
		next = startNode
	)

	for i := 0; i < len(visits); i++ {
		visits[i] = -1
	}

	for {
		visits[next] = step
		step += 1
		if next == 0 {
			break
		}
		next = parents[next]
		if next < 0 {
			break
		}
	}
}

type solver struct {
	rewards   []int
	children  [][]int
	bobVisits []int
}

func (s solver) solve(aliceStartNode int) int {
	return s.iterate(aliceStartNode, 0)
}

func (s solver) iterate(node int, step int) int {
	var (
		bobVisitStep = s.bobVisits[node]
		children     = s.children[node]
		reward       = s.rewards[node]
		actualReward int
	)

	if bobVisitStep == -1 {
		actualReward = reward
	} else if step == bobVisitStep {
		actualReward = reward / 2
	} else if step < bobVisitStep {
		actualReward = reward
	} else {
		actualReward = 0
	}

	if len(children) == 0 {
		return actualReward
	}

	maxSubReward := math.MinInt32
	for _, childNode := range s.children[node] {
		subReward := s.iterate(childNode, step+1)
		if subReward > maxSubReward {
			maxSubReward = subReward
		}
	}

	return actualReward + maxSubReward
}
