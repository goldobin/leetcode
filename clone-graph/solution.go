package clone_graph

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	var (
		m = map[*Node]*Node{}
		q = []*Node{node}
	)

	for len(q) > 0 {
		orig := q[0]
		q = q[1:]

		if _, ok := m[orig]; ok {
			continue
		}

		cloned := Node{Val: orig.Val}
		m[orig] = &cloned
		if orig.Neighbors == nil {
			continue
		}
		cloned.Neighbors = make([]*Node, len(orig.Neighbors))
		q = append(q, orig.Neighbors...)
	}

	for orig, cloned := range m {
		for i, neighbor := range orig.Neighbors {
			cloned.Neighbors[i] = m[neighbor]
		}
	}

	return m[node]
}

type Node struct {
	Val       int
	Neighbors []*Node
}
