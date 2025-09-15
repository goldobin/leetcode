package game_of_life

import "fmt"

func gameOfLife(b1 [][]int) {
	if len(b1) == 0 {
		return
	}

	if len(b1[0]) == 0 {
		return
	}

	var (
		d  = point{len(b1), len(b1[0])}
		b2 = newBoard(d)
	)

	for i := 0; i < d.x; i++ {
		for j := 0; j < d.y; j++ {
			p := point{i, j}
			step(b2, b1, d, p)
		}
	}

	copy(b1, b2)
}

type point struct {
	x, y int
}

const (
	alive = 1
	dead  = 0
)

func newBoard(d point) [][]int {
	b := make([][]int, d.x)
	for i := 0; i < d.x; i++ {
		b[i] = make([]int, d.y)
	}
	return b
}

func step(dst, src [][]int, d, p point) {
	state := src[p.x][p.y]
	aliveCount := 0

	if p.x > 0 && src[p.x-1][p.y] == alive {
		aliveCount++
	}
	if p.x < d.x-1 && src[p.x+1][p.y] == alive {
		aliveCount++
	}
	if p.y > 0 && src[p.x][p.y-1] == alive {
		aliveCount++
	}
	if p.y < d.y-1 && src[p.x][p.y+1] == alive {
		aliveCount++
	}
	if p.x > 0 && p.y > 0 && src[p.x-1][p.y-1] == alive {
		aliveCount++
	}
	if p.x > 0 && p.y < d.y-1 && src[p.x-1][p.y+1] == alive {
		aliveCount++
	}
	if p.x < d.x-1 && p.y > 0 && src[p.x+1][p.y-1] == alive {
		aliveCount++
	}
	if p.x < d.x-1 && p.y < d.y-1 && src[p.x+1][p.y+1] == alive {
		aliveCount++
	}

	switch state {
	case alive:
		if aliveCount < 2 || aliveCount > 3 {
			state = dead
		}
	case dead:
		if aliveCount == 3 {
			state = alive
		}
	default:
		panic(fmt.Sprintf("invalid state %d", state))
	}

	dst[p.x][p.y] = state
}
