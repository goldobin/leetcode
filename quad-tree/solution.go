package quad_tree

func construct(grid [][]int) *Node {
	return node(grid, 0, 0, len(grid))
}

func node(grid [][]int, row, col, length int) *Node {
	if length == 1 {
		return &Node{
			Val:    grid[row][col] == 1,
			IsLeaf: true,
		}
	}

	if length == 2 {
		var (
			r1  = row + 1
			c1  = col + 1
			v00 = grid[row][col] == 1
			v01 = grid[row][c1] == 1
			v10 = grid[r1][col] == 1
			v11 = grid[r1][c1] == 1
		)
		if isZero := !v00 && !v01 && !v10 && !v11; isZero {
			return &Node{
				Val: false, IsLeaf: true,
			}
		}
		if isOne := v00 && v01 && v10 && v11; isOne {
			return &Node{
				Val: true, IsLeaf: true,
			}
		}

		return &Node{
			Val:         false,
			IsLeaf:      false,
			TopLeft:     &Node{Val: v00, IsLeaf: true},
			TopRight:    &Node{Val: v01, IsLeaf: true},
			BottomLeft:  &Node{Val: v10, IsLeaf: true},
			BottomRight: &Node{Val: v11, IsLeaf: true},
		}
	}

	var (
		hl  = length / 2
		r1  = row + hl
		c1  = col + hl
		n00 = node(grid, row, col, hl)
		n01 = node(grid, row, c1, hl)
		n10 = node(grid, r1, col, hl)
		n11 = node(grid, r1, c1, hl)
	)

	if isLeaf := n00.IsLeaf && n01.IsLeaf && n10.IsLeaf && n11.IsLeaf; isLeaf {
		var (
			v00 = n00.Val
			v01 = n01.Val
			v10 = n10.Val
			v11 = n11.Val
		)
		if isZero := !v00 && !v01 && !v10 && !v11; isZero {
			return &Node{
				Val: false, IsLeaf: true,
			}
		}
		if isOne := v00 && v01 && v10 && v11; isOne {
			return &Node{
				Val: true, IsLeaf: true,
			}
		}
	}

	return &Node{
		Val:         false,
		IsLeaf:      false,
		TopLeft:     n00,
		TopRight:    n01,
		BottomLeft:  n10,
		BottomRight: n11,
	}
}
