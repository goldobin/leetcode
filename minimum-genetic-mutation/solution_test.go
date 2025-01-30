package minimum_genetic_mutation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minMutation(t *testing.T) {
	tests := []struct {
		name      string
		startGene string
		endGene   string
		bank      []string
		want      int
	}{
		{
			name:      "case 1.1 leetcode",
			startGene: "AACCGGTT",
			endGene:   "AACCGGTA",
			bank:      []string{"AACCGGTA"},
			want:      1,
		},
		{
			name:      "case 1.2 leetcode",
			startGene: "AACCGGTT",
			endGene:   "AAACGGTA",
			bank:      []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"},
			want:      2,
		},
		{
			name:      "case 1.3 leetcode",
			startGene: "AACCGGTT",
			endGene:   "AACCGGTA",
			bank:      []string{},
			want:      -1,
		},
		{
			name:      "case 1.4 leetcode",
			startGene: "AAAAAAAA",
			endGene:   "CCCCCCCC",
			bank: []string{
				"AAAAAAAA",
				"AAAAAAAC",
				"AAAAAACC",
				"AAAAACCC",
				"AAAACCCC",
				"AACACCCC",
				"ACCACCCC",
				"ACCCCCCC",
				"CCCCCCCA",
			},
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minMutation(tt.startGene, tt.endGene, tt.bank)
			assert.Equal(t, tt.want, got)
		})
	}
}
