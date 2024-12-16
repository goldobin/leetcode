package linked_list_loop_detection

type node struct {
	value int
	next  *node
}

func findCycle(head *node) *node {
	seen := make(map[*node]struct{})
	current := head

	for {
		if current == nil {
			return nil
		}

		if _, ok := seen[current]; ok {
			return current
		}
		seen[current] = struct{}{}
		current = current.next
	}
}
