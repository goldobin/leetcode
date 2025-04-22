package find_all_possible_recipes_from_given_supplies

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findAllRecipes(t *testing.T) {
	tests := []struct {
		name        string
		recipes     []string
		ingredients [][]string
		supplies    []string
		want        []string
	}{
		{
			name: "case 0.1 no recipes",
		},
		{
			name:        "case 1.1 leetcode",
			recipes:     []string{"bread"},
			ingredients: [][]string{{"yeast", "flour"}},
			supplies:    []string{"yeast", "flour", "corn"},
			want:        []string{"bread"},
		},
		{
			name:        "case 1.2 leetcode",
			recipes:     []string{"bread", "sandwich"},
			ingredients: [][]string{{"yeast", "flour"}, {"bread", "meat"}},
			supplies:    []string{"yeast", "flour", "meat"},
			want:        []string{"bread", "sandwich"},
		},
		{
			name:        "case 1.3 leetcode",
			recipes:     []string{"bread", "sandwich", "burger"},
			ingredients: [][]string{{"yeast", "flour"}, {"bread", "meat"}, {"sandwich", "meat", "bread"}},
			supplies:    []string{"yeast", "flour", "meat"},
			want:        []string{"bread", "sandwich", "burger"},
		},
		{
			name:    "case 1.4 leetcode",
			recipes: []string{"xevvq", "izcad", "p", "we", "bxgnm", "vpio", "i", "hjvu", "igi", "anp", "tokfq", "z", "kwdmb", "g", "qb", "q", "b", "hthy"},
			ingredients: [][]string{
				{"wbjr"},
				{"otr", "fzr", "g"},
				{"fzr", "wi", "otr", "xgp", "wbjr", "igi", "b"},
				{"fzr", "xgp", "wi", "otr", "tokfq", "izcad", "igi", "xevvq", "i", "anp"},
				{"wi", "xgp", "wbjr"},
				{"wbjr", "bxgnm", "i", "b", "hjvu", "izcad", "igi", "z", "g"},
				{"xgp", "otr", "wbjr"},
				{"wbjr", "otr"},
				{"wbjr", "otr", "fzr", "wi", "xgp", "hjvu", "tokfq", "z", "kwdmb"},
				{"xgp", "wi", "wbjr", "bxgnm", "izcad", "p", "xevvq"},
				{"bxgnm"},
				{"wi", "fzr", "otr", "wbjr"},
				{"wbjr", "wi", "fzr", "xgp", "otr", "g", "b", "p"},
				{"otr", "fzr", "xgp", "wbjr"},
				{"xgp", "wbjr", "q", "vpio", "tokfq", "we"},
				{"wbjr", "wi", "xgp", "we"},
				{"wbjr"},
				{"wi"},
			},
			supplies: []string{"wi", "otr", "wbjr", "fzr", "xgp"},
			want:     []string{"xevvq", "izcad", "bxgnm", "i", "hjvu", "tokfq", "z", "g", "b", "hthy"},
		},
		{
			name:    "case 2.1",
			recipes: []string{"bread", "sandwich", "burger", "chicken_burger"},
			ingredients: [][]string{
				{"yeast", "flour"},
				{"bread", "meat"},
				{"sandwich", "meat", "bread"},
				{"chicken", "bread"},
			},
			supplies: []string{"yeast", "flour", "meat"},
			want:     []string{"bread", "sandwich", "burger"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findAllRecipes(tt.recipes, tt.ingredients, tt.supplies)
			slices.Sort(tt.want)
			slices.Sort(got)
			assert.Equal(t, tt.want, got)
		})
	}
}
