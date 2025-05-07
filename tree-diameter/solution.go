package tree_diameter

func treeDiameter(edges [][]int) int {
	if len(edges) == 0 {
		return 0
	}

	var (
		t     = newTree(edges)
		queue = t.collectLeafs()
	)

	if len(queue) < 2 {
		panic("tree should have at least two leaf nodes")
	}

	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]

		if len(n.edges) != 1 {
			panic("only the nodes with precisely one edge is expected to be in the queue")
		}

		var adjN *node
		for k := range n.edges {
			adjN = k
		}

		delete(adjN.edges, n)
		if len(adjN.edges) == 0 {
			// We have met the opposite walker
			return n.shoulderLen + adjN.shoulderLen + 1
		}

		adjN.shoulderLen = max(n.shoulderLen+1, adjN.shoulderLen)
		if len(adjN.edges) == 1 {
			queue = append(queue, adjN)
		}
	}

	return 0
}

type node struct {
	edges       map[*node]struct{}
	shoulderLen int
}

type tree struct {
	nodes map[int]*node
}

func newTree(edges [][]int) *tree {
	result := tree{
		nodes: make(map[int]*node),
	}

	for _, edge := range edges {
		var (
			n1 = result.getOrCreateNode(edge[0])
			n2 = result.getOrCreateNode(edge[1])
		)

		n1.edges[n2] = struct{}{}
		n2.edges[n1] = struct{}{}
	}

	return &result
}

func (t *tree) getOrCreateNode(i int) *node {
	n, ok := t.nodes[i]
	if !ok {
		n = &node{
			edges: make(map[*node]struct{}),
		}
		t.nodes[i] = n
	}
	return n
}

func (t *tree) collectLeafs() []*node {
	var result []*node
	for _, n := range t.nodes {
		if len(n.edges) != 1 {
			continue
		}
		result = append(result, n)
	}

	return result
}
