package find_the_index_of_the_first_occurrence_in_a_string

func strStr(haystack string, needle string) int {
	first := needle[0]
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if haystack[i] != first {
			continue
		}

		if len(needle) == 1 {
			return i
		}

		for j := 1; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
			if j == len(needle)-1 {
				return i
			}
		}
	}

	return -1
}
