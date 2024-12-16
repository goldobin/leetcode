package valid_sudoku

func isValidSudoku(board [][]byte) bool {
	if len(board) != 9 {
		return false
	}

	if len(board[0]) != 9 {
		return false
	}

	rows := make([]box, 9)
	cols := make([]box, 9)
	boxes := make([]box, 9)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if empty == board[i][j] {
				continue
			}

			v := board[i][j]
			if !put(rows, i, v) {
				return false
			}
			if !put(cols, j, v) {
				return false
			}
			if !put(boxes, boxIndex(i, j), v) {
				return false
			}
		}
	}

	return true
}

const empty = '.'

type box struct {
	nums map[byte]struct{}
}

func boxIndex(i, j int) int {
	c := j / 3
	r := i / 3
	return r*3 + c
}

func put(m []box, ix int, v byte) bool {
	b := &m[ix]
	if b.nums == nil {
		b.nums = make(map[byte]struct{}, 9)
	}

	if _, ok := b.nums[v]; ok {
		return false
	}
	b.nums[v] = struct{}{}
	return true
}
