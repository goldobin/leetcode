package longest_consecutive_sequence

import (
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestConsecutive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "case 0.1",
			nums: []int{},
			want: 0,
		},
		{
			name: "case 0.2",
			nums: []int{1, 2, 3, 4, 5},
			want: 5,
		},
		{
			name: "case 0.3",
			nums: []int{5, 4, 3, 2, 1},
			want: 5,
		},
		{
			name: "case 1 leetcode",
			nums: []int{100, 4, 200, 1, 3, 2},
			want: 4,
		},
		{
			name: "case 2 leetcode",
			nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			want: 9,
		},
		{
			name: "case 3 leetcode",
			nums: []int{-7, -1, 3, -9, -4, 7, -3, 2, 4, 9, 4, -9, 8, -7, 5, -1, -7},
			want: 4,
		},
		{
			name: "case 4 leetcode",
			nums: []int{-6, 8, -5, 7, -9, -1, -7, -6, -9, -7, 5, 7, -1, -8, -8, -2, 0},
			want: 5,
		},
		{
			name: "case 5 leetcode",
			nums: mustLoadData("data/case05.txt"),
			want: 25001,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestConsecutive(tt.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}

func loadData(fileName string) ([]int, error) {
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

	ss := strings.Split(string(bs), ",")
	ns := make([]int, len(ss))

	for i, s := range ss {
		ns[i], err = strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			return nil, fmt.Errorf("failed to convert string to int at position %v: %w", i, err)
		}
	}

	return ns, nil
}

func mustLoadData(fileName string) []int {
	ints, err := loadData(fileName)
	if err != nil {
		panic(err)
	}
	return ints
}
