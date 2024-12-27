package rotate_image

import "fmt"

func rotate(matrix [][]int) {
	radius := len(matrix) / 2
	l := len(matrix)

	for d := 0; d < radius; d++ {
		fmt.Printf("level: %d\n", d)
		for k := d; k < l-d-1; k++ {
			iTop := k
			jTop := d
			iBottom := l - k - 1
			jBottom := l - d - 1

			iLeft := l - d - 1
			jLeft := k

			iRight := d
			jRight := l - k - 1

			v := matrix[iTop][jTop]
			matrix[iTop][jTop] = matrix[iLeft][jLeft]
			matrix[iLeft][jLeft] = matrix[iBottom][jBottom]
			matrix[iBottom][jBottom] = matrix[iRight][jRight]
			matrix[iRight][jRight] = v
		}
	}
}
