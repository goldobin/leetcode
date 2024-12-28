package edit_distance_recursive

func minDistance(word1, word2 string) int {
	return iterate(word1, word2, len(word1), len(word2))
}

const (
	matchCostIx  = 0
	insertCostIx = 1
	deleteCostIx = 2
)

func iterate(s, t string, i, j int) int {
	costs := make([]int, 3)

	if i == 0 {
		return j * insertDeleteCost(' ')
	}

	if j == 0 {
		return i * insertDeleteCost(' ')
	}

	costs[matchCostIx] = iterate(s, t, i-1, j-1) + matchCost(s[i-1], t[j-1])
	costs[insertCostIx] = iterate(s, t, i, j-1) + insertDeleteCost(t[j-1])
	costs[deleteCostIx] = iterate(s, t, i-1, j) + insertDeleteCost(s[i-1])

	result := costs[matchCostIx]
	for _, c := range costs {
		if c < result {
			result = c
		}
	}

	return result
}

func matchCost(s, t byte) int {
	if s == t {
		return 0
	}

	return 1
}

func insertDeleteCost(_ byte) int {
	return 1
}
