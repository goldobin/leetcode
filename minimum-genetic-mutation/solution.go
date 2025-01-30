package minimum_genetic_mutation

import "math"

func minMutation(startGene string, endGene string, bank []string) int {
	if len(bank) == 0 {
		return -1
	}

	var (
		n       = len(bank) + 1
		startIx = len(bank)
		endIx   = -1
		g       = graph{
			vs: make([]vertex, n),
			es: make([][]edge, n),
		}
	)

	g.vs[startIx].val = startGene
	for i, gene := range bank {
		g.vs[i].val = gene
		if gene == endGene {
			endIx = i
		}
	}

	if endIx == -1 {
		return -1
	}

	for i := 0; i < len(g.vs); i++ {
		for j := i + 1; j < len(g.vs); j++ {
			if isNextMutation(g.vs[i].val, g.vs[j].val) {
				g.es[i] = append(g.es[i], edge{to: j, weight: 1})
				g.es[j] = append(g.es[j], edge{to: i, weight: 1})
			}
		}
	}

	if len(g.es[endIx]) == 0 {
		return -1
	}

	ds := make([]data, len(g.vs))
	dijkstra(ds, g, startIx)
	return ds[endIx].dist
}

func isNextMutation(gene1, gene2 string) bool {
	diff := 0
	for i := 0; i < len(gene1); i++ {
		if gene1[i] != gene2[i] {
			diff++
		}
	}
	return diff == 1
}

func dijkstra(dst []data, g graph, start int) {
	if len(dst) != len(g.vs) {
		panic("invalid input")
	}

	for i := 0; i < len(dst); i++ {
		dst[i].inTree = false
		dst[i].dist = math.MaxInt
		dst[i].parent = -1
	}

	dst[start].dist = 0
	v := start
	for !dst[v].inTree {
		dst[v].inTree = true

		for _, e := range g.es[v] {
			w := e.to
			newDist := dst[v].dist + e.weight
			if dst[w].dist <= newDist {
				continue
			}
			dst[w].dist = newDist
			dst[w].parent = v
		}

		dist := math.MaxInt
		for i := 0; i < len(dst); i++ {
			if dst[i].inTree || dst[i].dist >= dist {
				continue
			}
			dist = dst[i].dist
			v = i
		}
	}
}

type graph struct {
	vs []vertex
	es [][]edge
}

type vertex struct {
	val string
}

type edge struct {
	to     int
	weight int
}

type data struct {
	inTree bool
	dist   int
	parent int
}
