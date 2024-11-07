package edit_distance_recursive

func matchCost(s, t byte) int {
	if s == t {
		return 0
	}

	return 1
}

func insertDeleteCost(_ byte) int {
	return 1
}

func stringCompare(s, t string) int {
	return iterate(s, t, len(s), len(t))
}

const (
	matchCostIx  = 0
	insertCostIx = 1
	deleteCostIx = 2
)

func iterate(s, t string, i, j int) int {
	var (
		costs = make([]int, 3)
	)

	if i == 0 {
		return j * insertDeleteCost(' ')
	}

	if j == 0 {
		return i * insertDeleteCost(' ')
	}

	costs[matchCostIx] = iterate(s, t, i-1, j-1) + matchCost(s[i-1], t[j-1])
	costs[insertCostIx] = iterate(s, t, i, j-1) + insertDeleteCost(t[j-1])
	costs[deleteCostIx] = iterate(s, t, i-1, j) + insertDeleteCost(s[i-1])

	lowestCost := costs[matchCostIx]
	for _, c := range costs {
		if c < lowestCost {
			lowestCost = c
		}
	}

	return lowestCost
}
