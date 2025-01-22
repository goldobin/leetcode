package set_matrix_zeroes

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}

	if len(matrix[0]) == 0 {
		return
	}

	var (
		m    = len(matrix)
		n    = len(matrix[0])
		rows = make(map[int]bool, len(matrix))
		cols = make(map[int]bool, len(matrix[0]))
	)

	for i := 0; i < m; i++ {
		if len(matrix[i]) != len(matrix[0]) {
			return
		}

		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}

	for i, ok := range rows {
		if !ok {
			continue
		}
		for j := 0; j < n; j++ {
			matrix[i][j] = 0
		}
	}

	for j, ok := range cols {
		if !ok {
			continue
		}
		for i := 0; i < m; i++ {
			matrix[i][j] = 0
		}
	}
}
