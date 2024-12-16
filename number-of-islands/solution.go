package number_of_islands

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	if len(grid[0]) == 0 {
		return 0
	}

	var (
		d = point{len(grid), len(grid[0])}
		t = traverser{
			g: grid,
			v: make(visited, d.x),
			q: make(pointSet),
			d: d,
		}
	)

	for i := 0; i < d.x; i++ {
		t.v[i] = make([]bool, d.y)
	}

	t.start(point{0, 0})
	return t.i
}

type (
	point    struct{ x, y int }
	pointSet map[point]struct{}
	grid     [][]byte
	visited  [][]bool

	traverser struct {
		g grid
		v visited
		q pointSet
		d point
		i int
	}
)

const (
	land = '1'
)

func (t *traverser) start(p point) {
	t.enqueue(p)

	for {
		np, ok := t.dequeue()
		if !ok {
			break
		}
		k := t.kind(np)
		if k == land {
			t.i += 1
		}

		t.traverse(k, np)
	}
}

func (t *traverser) enqueue(p point) {
	if t.visited(p) {
		return
	}
	t.q[p] = struct{}{}
}

func (t *traverser) dequeue() (point, bool) {
	for e := range t.q {
		delete(t.q, e)
		return e, true
	}

	return point{}, false
}

func (t *traverser) kind(p point) byte {
	return t.g[p.x][p.y]
}

func (t *traverser) visited(p point) bool {
	return t.v[p.x][p.y]
}

func (t *traverser) markAsVisited(p point) {
	if t.visited(p) {
		return
	}

	t.v[p.x][p.y] = true
	delete(t.q, p)
}

func (t *traverser) traverse(kind byte, p point) {
	if t.visited(p) {
		return
	}

	t.markAsVisited(p)

	var (
		neighbors = make([]point, 0, 4)
	)

	if p.x > 0 {
		np := point{p.x - 1, p.y}
		if !t.visited(np) {
			neighbors = append(neighbors, np)
		}
	}
	if p.x < t.d.x-1 {
		np := point{p.x + 1, p.y}
		if !t.visited(np) {
			neighbors = append(neighbors, np)
		}
	}
	if p.y > 0 {
		np := point{p.x, p.y - 1}
		if !t.visited(np) {
			neighbors = append(neighbors, np)
		}
	}
	if p.y < t.d.y-1 {
		np := point{p.x, p.y + 1}
		if !t.visited(np) {
			neighbors = append(neighbors, np)
		}
	}

	for _, np := range neighbors {
		if t.kind(np) != kind {
			t.enqueue(np)
			continue
		}

		t.traverse(kind, np)
	}
}
