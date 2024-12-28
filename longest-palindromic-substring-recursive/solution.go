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

	if s[0] == s[len(s)-1] {
		p := longestPalindrome(s[1 : len(s)-1])
		if len(p) == len(s)-2 {
			return s
		}

		return p
	}

	p1 := longestPalindrome(s[:len(s)-1])
	p2 := longestPalindrome(s[1:])

	if len(p1) <= len(p2) {
		return p2
	}

	return p1
}
