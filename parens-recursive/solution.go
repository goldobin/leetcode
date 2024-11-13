package parens_recursive

import (
	"maps"
	"slices"
)

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}

	if n == 1 {
		return []string{"()"}
	}

	result := make(map[string]struct{})

	ss11 := generateParenthesis(n - 1)
	for _, s11 := range ss11 {
		result["("+s11+")"] = struct{}{}
	}

	for i := 1; i < n; i++ {
		ss21 := generateParenthesis(i)
		ss22 := generateParenthesis(n - i)
		for _, s21 := range ss21 {
			for _, s22 := range ss22 {
				result[s21+s22] = struct{}{}
			}
		}
	}

	return slices.Collect(maps.Keys(result))
}
