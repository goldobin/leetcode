package triangle_dynamic

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}

	for i := len(triangle) - 2; i >= 0; i-- {
		row := triangle[i]
		for j := 0; j < len(row); j++ {
			v := row[j]
			nextRowI := i + 1
			left := triangle[nextRowI][j+1]
			right := triangle[nextRowI][j]
			if left < right {
				row[j] = v + left
			} else {
				row[j] = v + right
			}
		}
	}

	return triangle[0][0]
}
