package integer_to_roman

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_intToRoman(t *testing.T) {
	tests := []struct {
		in   int
		want string
	}{
		{1, "I"},
		{4, "IV"},
		{3, "III"},
		{58, "LVIII"},
		{1994, "MCMXCIV"},
		{3749, "MMMDCCXLIX"},
		{3999, "MMMCMXCIX"},
		{4000, ""},
		{0, ""},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("%v -> %v", tt.in, tt.want)
		t.Run(name, func(t *testing.T) {
			got := intToRoman(tt.in)
			assert.Equal(t, tt.want, got)
		})
	}
}
