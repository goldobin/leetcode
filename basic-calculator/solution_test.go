package basic_calculator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Calculate(t *testing.T) {
	t.Skip("solution not fully implemented")
	tests := []struct {
		in   string
		want int
	}{
		{"-1", -1},
		{"0 + 0", 0},
		{"0 + 1", 1},
		{"1 + 0", 1},
		{"1 + 1", 2},
		{"1 - 1", 0},
		{"1 - 2", -1},
		{"10 - 2", 8},
		{"10 - 20", -10},
		{"(10+11) - 20", 1},
		{"-(10+11) - 20", -41},
		{"5 + ((10+11) - 20)", 6},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got := calculate(tt.in)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_parseNextToken(t *testing.T) {
	tests := []struct {
		in         string
		wantToken  token
		wantOffset int
		wantErr    bool
	}{
		{in: "+", wantToken: plusToken, wantOffset: 1},
		{in: " +", wantToken: whitespaceToken, wantOffset: 1},
		{in: "-", wantToken: minusToken, wantOffset: 1},
		{in: "(", wantToken: openGroupToken, wantOffset: 1},
		{in: "  (  ", wantToken: whitespaceToken, wantOffset: 2},
		{in: ")", wantToken: closeGroupToken, wantOffset: 1},
		{in: ") ", wantToken: closeGroupToken, wantOffset: 1},
		{in: " ", wantToken: whitespaceToken, wantOffset: 1},
		{in: "    ", wantToken: whitespaceToken, wantOffset: 4},
		{in: "0", wantToken: numberToken(0), wantOffset: 1},
		{in: "1", wantToken: numberToken(1), wantOffset: 1},
		{in: "9", wantToken: numberToken(9), wantOffset: 1},
		{in: "0", wantToken: numberToken(0), wantOffset: 1},
		{in: "100", wantToken: numberToken(100), wantOffset: 3},
		{in: "   100", wantToken: whitespaceToken, wantOffset: 3},
		{in: "09", wantToken: numberToken(9), wantOffset: 2},
		{in: "09x", wantToken: numberToken(9), wantOffset: 2},
		{in: "  09x", wantToken: whitespaceToken, wantOffset: 2},
		{in: "x", wantErr: true},
		{in: "x ", wantErr: true},
		{in: "x1", wantErr: true},
		{in: "  x", wantToken: whitespaceToken, wantOffset: 2},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			gotOffset, gotToken, gotErr := parseNextToken(tt.in)
			if tt.wantErr {
				assert.Error(t, gotErr)
				return
			}

			if !assert.NoError(t, gotErr) {
				return
			}

			if !assert.LessOrEqual(t, gotOffset, len(tt.in)) {
				return
			}
			assert.Equal(t, tt.wantOffset, gotOffset)
			assert.Equal(t, tt.wantToken, gotToken)
		})
	}
}

func Test_parse(t *testing.T) {
	tests := []struct {
		in         string
		wantTokens []token
		wantErr    bool
	}{
		{in: "", wantErr: true},
		{in: " ", wantTokens: []token{}},
		{in: "   ", wantTokens: []token{}},
		{in: " ( ) ", wantTokens: []token{openGroupToken, closeGroupToken}},
		{in: "1 + 1", wantTokens: []token{numberToken(1), plusToken, numberToken(1)}},
		{in: "1 * 1", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			gotTokens, gotErr := parse(tt.in)

			if tt.wantErr {
				assert.Error(t, gotErr)
				return
			}

			if !assert.NoError(t, gotErr) {
				return
			}

			assert.Equal(t, tt.wantTokens, gotTokens)
		})
	}
}
