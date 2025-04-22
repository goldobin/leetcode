package find_all_possible_recipes_from_given_supplies

type node struct {
	name        string
	recipes     []*node
	ingredients []*node
	hitCount    int
}

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	var (
		nodes  = make(map[string]*node, len(recipes)+len(ingredients))
		queue  []*node
		result []string
	)

	getOrCreateNode := func(name string) *node {
		n, ok := nodes[name]
		if ok {
			return n
		}

		n = &node{
			name: name,
		}
		nodes[name] = n
		return n
	}

	for rIx, recipe := range recipes {
		rNode := getOrCreateNode(recipe)
		for _, ingredient := range ingredients[rIx] {
			iNode := getOrCreateNode(ingredient)
			iNode.recipes = append(iNode.recipes, rNode)
			rNode.ingredients = append(rNode.ingredients, iNode)
		}
	}

	for _, s := range supplies {
		n, ok := nodes[s]
		if !ok {
			continue
		}

		queue = append(queue, n)
	}

	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]

		for _, r := range n.recipes {
			r.hitCount += 1
			if r.hitCount >= len(r.ingredients) {
				queue = append(queue, r)
				result = append(result, r.name)
			}
		}
	}

	return result
}
