package parens_dynamic

import (
	"maps"
	"slices"
)

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}

	memory := make([][]string, n+1)
	memory[0] = []string{}
	memory[1] = []string{"()"}

	for i := 2; i <= n; i++ {
		result := make(map[string]struct{})

		ss11 := memory[i-1]
		for _, s11 := range ss11 {
			result["("+s11+")"] = struct{}{}
		}

		for j := 1; j < i; j++ {
			ss21 := memory[j]
			ss22 := memory[i-j]
			for _, s21 := range ss21 {
				for _, s22 := range ss22 {
					result[s21+s22] = struct{}{}
				}
			}
		}

		memory[i] = slices.Collect(maps.Keys(result))
	}

	return memory[n]
}
