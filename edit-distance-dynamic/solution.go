package edit_distance_dynamic

import (
	"slices"
	"strings"
)

type result struct {
	cost     int
	editPath []int
}

func stringCompare(s, t string) result {
	var (
		size    = maxLen(s, t)
		m       = makeSquareMatrix(size + 1)
		paddedS = leftPad(s, 1)
		paddedT = leftPad(t, 1)
	)

	initFirstRow(m)
	initFirstCol(m)

	return result{
		cost:     run(paddedS, paddedT, m),
		editPath: editPath(paddedS, paddedT, m),
	}
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

func leftPad(s string, size int) string {
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
)

func matchCost(s, t byte) int {
	if s == t {
		return 0
	}

	return 1
}

func insertDeleteCost(_ byte) int {
	return 1
}

func run(s, t string, m [][]cell) int {
	costs := make([]int, 3)

	// Sting is left padded, so we start from 1
	for i := 1; i < len(s); i++ {
		for j := 1; j < len(t); j++ {
			costs[matchOp] = m[i-1][j-1].cost + matchCost(s[i], t[j])
			costs[insertOp] = m[i][j-1].cost + insertDeleteCost(t[j])
			costs[deleteOp] = m[i-1][j].cost + insertDeleteCost(s[i])

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

	return m[len(s)-1][len(t)-1].cost
}

func editPath(s, t string, m [][]cell) []int {
	return editPathIterate(s, t, m, len(s)-1, len(t)-1)
}

func editPathIterate(s, t string, m [][]cell, i, j int) []int {
	var (
		p       = m[i][j].parent
		prevOps []int
	)

	switch p {
	case -1:
		return []int{}
	case matchOp:
		prevOps = editPathIterate(s, t, m, i-1, j-1)
	case insertOp:
		prevOps = editPathIterate(s, t, m, i, j-1)
	case deleteOp:
		prevOps = editPathIterate(s, t, m, i-1, j)
	default:
		panic("unexpected parent value")
	}

	return slices.Concat(prevOps, []int{p})
}
