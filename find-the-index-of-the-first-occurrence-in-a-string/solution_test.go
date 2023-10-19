package find_the_index_of_the_first_occurrence_in_a_string

import (
	"fmt"
	"testing"
)

func Test_strStr(t *testing.T) {
	cases := []struct {
		haystack string
		needle   string
		expected int
	}{
		{
			haystack: "leetcode",
			needle:   "leeto",
			expected: -1,
		},
		{
			haystack: "sadbutsad",
			needle:   "sad",
			expected: 0,
		},
		{
			haystack: "sadbuttrue",
			needle:   "true",
			expected: 6,
		},
		{
			haystack: "sadbuttrueandother",
			needle:   "true",
			expected: 6,
		},
		{
			haystack: "sadbuttru",
			needle:   "true",
			expected: -1,
		},
		{
			haystack: "sadbuttrue",
			needle:   "t",
			expected: 5,
		},
		{
			haystack: "mississippi",
			needle:   "issip",
			expected: 4,
		},
	}

	for i, c := range cases {
		caseName := fmt.Sprintf("Case %02d", i+1)
		t.Run(caseName, func(t *testing.T) {
			result := strStr(c.haystack, c.needle)
			if result != c.expected {
				t.Errorf("expected=%d, got=%d", c.expected, result)
			}
		})
	}
}

func Benchmark_strStr(b *testing.B) {
	cases := []struct {
		haystack string
		needle   string
		expected int
	}{
		{
			haystack: "leetcode",
			needle:   "leeto",
			expected: -1,
		},
		{
			haystack: "sadbutsad",
			needle:   "sad",
			expected: 0,
		},
		{
			haystack: "sadbuttrue",
			needle:   "true",
			expected: 6,
		},
	}

	for _, c := range cases {
		strStr(c.haystack, c.needle)
	}
}
