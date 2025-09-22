package surrounded_regions

import "fmt"

func solve(board [][]byte) {
	height := len(board)
	if height < 3 {
		return
	}

	width := len(board[0])
	if width < 3 {
		return
	}

	// Pre-allocating queue and region slices. Those slices will be reused during algorithm execution
	// so no additional allocations needed
	s := solver{
		height:  height,
		width:   width,
		board:   board,
		visited: make([]bool, width*height),
		queue:   make([]pos, 0, width*height),
		region:  make([]pos, 0, width*height),
	}

	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			p1 := pos{row, col}
			v := board[row][col]
			if v != 'O' {
				continue
			}

			s.findRegion(p1)
			for _, p2 := range s.region {
				if !s.regionCanBeSurrounded {
					continue
				}
				board[p2.row][p2.col] = 'X'
			}
		}
	}
}

type (
	pos    struct{ row, col int }
	solver struct {
		height                int
		width                 int
		board                 [][]byte
		visited               []bool
		queue                 []pos
		regionCanBeSurrounded bool
		region                []pos
	}
)

func (s *solver) at(p pos) byte {
	s.mustBeOnBoard(p)
	return s.board[p.row][p.col]
}

func (s *solver) isOnBoard(p pos) bool {
	return p.row > -1 && p.col > -1 && p.row < s.height && p.col < s.width
}

func (s *solver) mustBeOnBoard(p pos) {
	if !s.isOnBoard(p) {
		panic(fmt.Sprintf("invalid coords: %v", p))
	}
}

func (s *solver) canBeSurrounded(p pos) bool {
	return p.row > 0 && p.col > 0 && p.row < s.height-1 && p.col < s.width-1
}

func (s *solver) isVisited(p pos) bool {
	s.mustBeOnBoard(p)
	return s.visited[p.row*s.width+p.col]
}

func (s *solver) markVisited(p pos, v bool) {
	s.mustBeOnBoard(p)
	s.visited[p.row*s.width+p.col] = v
}

func (s *solver) findRegion(p pos) {
	v := s.at(p)

	s.queue = append(s.queue, p)
	s.regionCanBeSurrounded = true
	s.region = s.region[:0]

	for len(s.queue) > 0 {
		curPos := s.queue[0]
		s.queue = s.queue[1:]

		curValue := s.at(curPos)
		if curValue != v {
			continue
		}

		if s.isVisited(curPos) {
			continue
		}

		s.markVisited(curPos, true)
		if !s.canBeSurrounded(curPos) {
			s.regionCanBeSurrounded = false
		}

		s.region = append(s.region, curPos)
		var candidates [4]pos
		for i := 0; i < 4; i++ {
			candidates[i] = curPos
		}

		candidates[0].row -= 1
		candidates[1].row += 1
		candidates[2].col -= 1
		candidates[3].col += 1

		for _, candidatePos := range candidates {
			if !s.isOnBoard(candidatePos) {
				continue
			}
			if !s.isOnBoard(candidatePos) {
				continue
			}
			s.queue = append(s.queue, candidatePos)
		}
	}
}
