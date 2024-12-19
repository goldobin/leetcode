package h_index

import "slices"

func hIndex(citations []int) int {
	slices.Sort(citations)
	slices.Reverse(citations)

	p := 1
	for _, cc := range citations {
		if cc == p {
			return p
		} else if cc < p {
			return p - 1
		}
		p += 1
	}

	return p - 1
}
