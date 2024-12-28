package edit_distance_dynamic

import (
	"strings"
)

func minDistance(word1, word2 string) int {
	var (
		size        = maxLen(word1, word2)
		m           = makeSquareMatrix(size + 1)
		paddedWord1 = padLeft(word1, 1)
		paddedWord2 = padLeft(word2, 1)
	)

	initFirstRow(m)
	initFirstCol(m)

	return iterate(paddedWord1, paddedWord2, m)
}

func maxLen(s, t string) int {
	if len(s) > len(t) {
		return len(s)
	}

	return len(t)
}

type cell struct {
	cost   int
	parent int
}

func makeSquareMatrix(size int) [][]cell {
	m := make([][]cell, size)
	for i := 0; i < size; i++ {
		m[i] = make([]cell, size)
	}

	return m
}

func initFirstRow(m [][]cell) {
	for i := 0; i < len(m); i++ {
		m[0][i].cost = i

		var p int
		if i == 0 {
			p = -1
		} else {
			p = insertOp
		}
		m[0][i].parent = p
	}
}

func initFirstCol(m [][]cell) {
	for i := 0; i < len(m); i++ {
		m[i][0].cost = i

		var p int
		if i == 0 {
			p = -1
		} else {
			p = deleteOp
		}
		m[i][0].parent = p
	}
}

func iterate(word1, word2 string, m [][]cell) int {
	costs := make([]int, 3)

	// Sting is left padded, so we start from 1
	for i := 1; i < len(word1); i++ {
		for j := 1; j < len(word2); j++ {
			var matchCost int
			if word1[i] != word2[j] {
				matchCost = differenceCost
			}

			costs[matchOp] = m[i-1][j-1].cost + matchCost
			costs[insertOp] = m[i][j-1].cost + insertDeleteCost
			costs[deleteOp] = m[i-1][j].cost + insertDeleteCost

			m[i][j].cost = costs[matchOp]
			m[i][j].parent = matchOp

			for k := insertOp; k <= deleteOp; k++ {
				if costs[k] < m[i][j].cost {
					m[i][j].cost = costs[k]
					m[i][j].parent = k
				}
			}
		}
	}

	return m[len(word1)-1][len(word2)-1].cost
}

func padLeft(s string, size int) string {
	if size <= 0 {
		return s
	}

	b := strings.Builder{}
	b.Grow(len(s) + size)

	for i := 0; i < size; i++ {
		b.WriteByte(' ')
	}

	b.Write([]byte(s))
	return b.String()
}

const (
	matchOp  = 0
	insertOp = 1
	deleteOp = 2

	differenceCost   = 1
	insertDeleteCost = 1
)
