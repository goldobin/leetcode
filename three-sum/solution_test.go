package three_sum

import (
	"fmt"
	"io"
	"os"
	"path"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_threeSum(t *testing.T) {
	tests := []struct {
		name          string
		nums          []int
		want          [][]int
		runBruteForce bool
	}{
		{
			name: "case 1.1 leetcode",
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "case 1.2 leetcode",
			nums: []int{0, 1, 1},
			want: [][]int{},
		},
		{
			name: "case 1.3 leetcode",
			nums: []int{0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
		{
			name:          "case 1.4 leetcode",
			nums:          mustParseNums(mustReadFile("data/nums04.txt")),
			want:          mustParseOut(mustReadFile("data/out04.txt")),
			runBruteForce: false,
		},
		{
			name: "case 1.5 leetcode",
			nums: []int{-2, 0, 1, 1, 2},
			want: [][]int{{-2, 0, 2}, {-2, 1, 1}},
		},
		{
			name: "case 1.6 leetcode",
			nums: []int{0, 0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortSlice2D(tt.want)

			if tt.runBruteForce {
				t.Run("brute force", func(t *testing.T) {
					got := threeSumBruteForce(tt.nums)
					assertTriadsExpected(t, tt.want, got)
				})
			}

			t.Run("fast", func(t *testing.T) {
				got := threeSum(tt.nums)
				assertTriadsExpected(t, tt.want, got)
			})
		})
	}
}

func readFile(fileName string) ([]byte, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get working directory:%w", err)
	}

	p := path.Join(dir, fileName)
	f, err := os.Open(p)
	if err != nil {
		return nil, fmt.Errorf("failed to open file for read %s: %w", p, err)
	}

	bs, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", p, err)
	}

	return bs, nil
}

func mustReadFile(fileName string) []byte {
	bs, err := readFile(fileName)
	if err != nil {
		panic(err)
	}
	return bs
}

func parseNums(bs []byte) ([]int, error) {
	var (
		ss = strings.Split(string(bs), ",")
		ns = make([]int, len(ss))
	)

	for i, s := range ss {
		var err error
		ns[i], err = strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			return nil, fmt.Errorf("failed to convert string to int at position %v: %w", i, err)
		}
	}

	return ns, nil
}

func mustParseNums(bs []byte) []int {
	ints, err := parseNums(bs)
	if err != nil {
		panic(err)
	}
	return ints
}

func parseOut(bs []byte) ([][]int, error) {
	var (
		re     = regexp.MustCompile(`\[(-?\d+),(-?\d+),(-?\d+)]`)
		result = make([][]int, 0)
		parts  = re.FindAll(bs, -1)
	)

	for _, part := range parts {
		subparts := re.FindSubmatch(part)
		if len(subparts) != 4 {
			return nil, fmt.Errorf("failed to parse out %v: expected 3 numbers was %v", string(part), len(subparts))
		}

		var (
			triplet = make([]int, 3)
			err     error
		)
		for i, subpart := range subparts[1:] {
			triplet[i], err = strconv.Atoi(string(subpart))
			if err != nil {
				return nil, fmt.Errorf("failed to parse out %v: %w", string(subpart), err)
			}
		}
		result = append(result, triplet)
	}

	return result, nil
}

func mustParseOut(bs []byte) [][]int {
	triplets, err := parseOut(bs)
	if err != nil {
		panic(err)
	}
	return triplets
}

func sortSlice2D(ts [][]int) {
	for i := 0; i < len(ts); i++ {
		slices.Sort(ts[i])
	}

	slices.SortFunc(ts, slices.Compare)
}

func assertTriadsSumsToZero(t *testing.T, got [][]int) bool {
	for i, v := range got {
		if !assert.Len(t, v, 3) {
			return false
		}
		s := v[0] + v[1] + v[2]
		if !assert.Equal(t, 0, s, nil, "got %v; want 0, at index %d", s, i) {
			return false
		}
	}

	return true
}

func compactSlice2D(got [][]int) [][]int {
	return slices.CompactFunc(got, slices.Equal)
}

func assertTriadsExpected(t *testing.T, expected, got [][]int) {
	sortSlice2D(got)
	if !assertTriadsSumsToZero(t, got) {
		return
	}

	compactedGot := compactSlice2D(slices.Clone(got))
	if !assert.Equal(t, compactedGot, got, "result contains duplicates") {
		return
	}

	assert.Equal(t, len(expected), len(got), "result contains wrong number of elements")
	assert.Equal(t, expected, got)
}

func Benchmark_threeSum(b *testing.B) {
	nums := mustParseNums(mustReadFile("data/nums04.txt"))

	for i := 0; i < b.N; i++ {
		got := threeSum(nums)
		if got == nil {
			b.Error("got nil; want non-nil")
		}
	}
}
