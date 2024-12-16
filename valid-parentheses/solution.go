package valid_parentheses

import "fmt"

func isValid(s string) bool {
	st := make(stack, 0, len(s))

	for i := 0; i < len(s); i++ {
		var v2 = s[i]
		if isLeft(v2) {
			st.push(v2)
			continue
		}

		if isRight(v2) {
			v1, ok := st.pull()
			if !ok {
				return false
			}
			if !doesMatch(v1, v2) {
				return false
			}
			continue
		}

		panic(fmt.Sprintf("unsupported character: %c", v2))
	}

	return len(st) == 0
}

type stack []uint8

func (s *stack) push(b uint8) {
	*s = append(*s, b)
}

func (s *stack) pull() (uint8, bool) {
	if len(*s) == 0 {
		return 0, false
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v, true
}

func isLeft(v uint8) bool {
	return v == '(' || v == '[' || v == '{'
}

func isRight(v uint8) bool {
	return v == ')' || v == ']' || v == '}'
}

func doesMatch(v1, v2 uint8) bool {
	return (v1 == '(' && v2 == ')') ||
		(v1 == '[' && v2 == ']') ||
		(v1 == '{' && v2 == '}')
}
