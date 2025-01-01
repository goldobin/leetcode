package longest_palindromic_substring_recursive

import "fmt"

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	if len(s) == 1 {
		return s
	}

	if len(s) == 2 {
		if s[0] == s[1] {
			return s
		}
		return s[0:1]
	}

	if len(s) == 3 {
		if s[0] == s[2] {
			return s
		}
	}

	var (
		maxC = maxCoords{v: 1}
		tr   = newTriangleFromString(s)
	)

	for i := 2; i < tr.sideLen(); i++ {
		for l := 0; l < tr.sideLen()-i; l++ {
			var (
				c            = coords{l, l + i}
				prevC        = coords{c.l + 1, c.r - 1}
				strC         = coords{c.l - 1, c.r - 1}
				d            = prevC.distance()
				v            = tr.locate(prevC)
				isPalindrome = len(v) == d &&
					strC.l >= 0 &&
					strC.r >= 0 &&
					strC.l < len(s) &&
					strC.r < len(s) &&
					s[strC.l] == s[strC.r]
			)

			maxC.update(prevC, len(v))

			if isPalindrome {
				v = s[strC.l : strC.r+1]
				maxC.update(c, len(v))
			}

			tr.update(c, v)
		}
	}

	result := tr.locate(maxC.coords)
	if result == "" {
		return s[0:1]
	}

	return result
}

func newTriangle(n int) triangle {
	tr := make([][]string, n)
	for i := 0; i < n; i++ {
		tr[i] = make([]string, i+1)
	}
	return tr
}

func newTriangleFromString(s string) triangle {
	tr := newTriangle(len(s) + 2)
	tr.fillZeroDiagonal(s)
	tr.fillFirstDiagonal(s)
	return tr
}

func (tr *triangle) fillZeroDiagonal(s string) {
	sideLen := len(*tr)
	if sideLen != len(s)+2 {
		panic(fmt.Sprintf("sting len=%d, got=%d", sideLen-2, len(s)))
	}

	for i := 0; i < len(s); i++ {
		(*tr)[i+1][i+1] = s[i : i+1]
	}
}

func (tr *triangle) fillFirstDiagonal(s string) {
	sideLen := len(*tr)
	if sideLen != len(s)+2 {
		panic(fmt.Sprintf("sting len=%d, got=%d", sideLen-2, len(s)))
	}

	for l := 0; l < len(s)-1; l++ {
		var (
			r = l + 1
			v string
		)
		if s[l] == s[r] {
			v = s[l : r+1]
		} else {
			v = s[l : l+1]
		}
		(*tr)[r+1][l+1] = v
	}
}

type (
	triangle [][]string
	coords   struct {
		l, r int
	}
	maxCoords struct {
		coords
		v int
	}
)

func (mtc *maxCoords) update(c coords, v int) {
	if v > mtc.v {
		mtc.l = c.l
		mtc.r = c.r
		mtc.v = v
	}
}

func (t coords) distance() int {
	return t.r - t.l + 1
}

func (tr *triangle) locate(c coords) string {
	return (*tr)[c.r][c.l]
}

func (tr *triangle) update(c coords, v string) {
	(*tr)[c.r][c.l] = v
}

func (tr *triangle) sideLen() int {
	return len(*tr)
}
