package majority_element

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_majorityElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "case 1",
			nums: []int{3, 2, 3},
			want: 3,
		},
		{
			name: "case 2",
			nums: []int{2, 2, 1, 1, 1, 2, 2},
			want: 2,
		},
		{
			name: "case 3",
			nums: []int{2, 2, 1, 1, 1, 1, 1, 2, 2},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := majorityElement(tt.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}
