package triangle_recursive

func minimumTotal(triangle [][]int) int {
	return iterate(triangle, 0, 0)
}

func iterate(triangle [][]int, i, j int) int {
	if i >= len(triangle) {
		return 0
	}

	v := triangle[i][j]
	left := iterate(triangle, i+1, j)
	right := iterate(triangle, i+1, j+1)

	if left < right {
		return v + left
	} else {
		return v + right
	}
}
