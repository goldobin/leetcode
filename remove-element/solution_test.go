package remove_element

import (
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

func Test_removeElement(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		val      int
		want     int
		wantNums []int
	}{
		{
			name:     "case 0.1",
			nums:     []int{},
			val:      2,
			want:     0,
			wantNums: []int{},
		},
		{
			name:     "case 0.2",
			nums:     []int{2},
			val:      2,
			want:     0,
			wantNums: []int{},
		},
		{
			name:     "case 0.3",
			nums:     []int{2, 2, 2, 2, 2},
			val:      2,
			want:     0,
			wantNums: []int{},
		},
		{
			name:     "case 0.4",
			nums:     []int{2, 2, 2, 2, 2},
			val:      3,
			want:     5,
			wantNums: []int{2, 2, 2, 2, 2},
		},
		{
			name:     "case 1 leetcode",
			nums:     []int{3, 2, 2, 3},
			val:      3,
			want:     2,
			wantNums: []int{2, 2},
		},
		{
			name:     "case 2 leetcode",
			nums:     []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:      2,
			want:     5,
			wantNums: []int{0, 1, 3, 0, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gotNums := slices.Clone(tt.nums)
			got := removeElement(gotNums, tt.val)

			if !assert.Equal(t, tt.want, got) {
				return
			}

			gotNums = gotNums[:got]
			slices.Sort(gotNums)
			slices.Sort(tt.wantNums)
			assert.Equal(t, tt.wantNums, gotNums)
		})
	}
}
