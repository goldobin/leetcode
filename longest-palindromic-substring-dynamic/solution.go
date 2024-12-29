package longest_palindromic_substring_recursive

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

	tr := makeTriangle(len(s))
	fillZeroDiagonal(tr, s)
	fillFirstDiagonal(tr, s)

	maxLen := 1
	maxL := 0
	maxR := 0

	for i := 2; i < len(s); i++ {
		for l := 0; l < len(s)-i; l++ {
			r := l + i
			prevL := l + 1
			prevR := r - 1
			d := prevR - prevL + 1
			v := tr[prevR][prevL]

			if len(v) > maxLen {
				maxLen = len(v)
				maxL = prevL
				maxR = prevR
			}

			if len(v) == d && s[l] == s[r] {
				v = s[l : r+1]
				if len(v) > maxLen {
					maxLen = len(v)
					maxL = l
					maxR = r
				}
			}

			tr[r][l] = v
		}
	}

	return tr[maxR][maxL]
}

func makeTriangle(n int) [][]string {
	tr := make([][]string, n)
	for i := 0; i < n; i++ {
		tr[i] = make([]string, i+1)
	}
	return tr
}

func fillZeroDiagonal(tr [][]string, s string) {
	for i := 0; i < len(s); i++ {
		tr[i][i] = s[i : i+1]
	}
}

func fillFirstDiagonal(tr [][]string, s string) {
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
		tr[r][l] = v
	}
}
