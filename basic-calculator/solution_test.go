package basic_calculator

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func Test_Calculate(t *testing.T) {
	//t.Skip("solution not fully implemented")
	tests := []struct {
		name string
		in   string
		want int
	}{
		{"unary negation", "-1", -1},
		{"zero sum", "0 + 0", 0},
		{"priority", " 2-1 + 2", 3},
		{"sum with zero r", "0 + 1", 1},
		{"sum with zero l", "1 + 0", 1},
		{"one plus one", "1 + 1", 2},
		{"one minus one", "1 - 1", 0},
		{"negative sum", "1 - 2", -1},
		{"case 01.1", "10 - 2", 8},
		{"case 01.2", "10 - 20", -10},
		{"case 01.3", "-10 - 20", -30},
		{"group case 01.1", "(10+11) - 20", 1},
		{"group case 01.2", "-(10+11) - 20", -41},
		{"group case 02.1", "(1+(4+5+2)-3)+(6+8)", 23},
		{"group case 02.2", "-(1+(4+5+2)-3)+(6+8)", 5},
		{"group case 02.3", "3-(1+(4+5+2)-3)+(6+8)", 8},
		{"group case 03", "5 + ((10+11) - 20)", 6},
		{"long case 01", string(mustReadFile("long_test_case_01.txt")), -1946},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
		{in: "1 + 1", wantTokens: []token{
			numberToken(1),
			plusToken,
			numberToken(1),
		}},
		{in: "1 * 1", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			gotTokens, gotErr := parseTokens(tt.in)

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

func mustReadFile(path string) []byte {

	// Get the directory of the current test file
	_, testFileName, _, ok := runtime.Caller(0)
	if !ok {
		panic("failed to get current test file directory")
	}
	dir := filepath.Dir(testFileName)

	// Construct the path to the target file
	filePath := filepath.Join(dir, path)

	// Open and read the file
	f, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Errorf("failed to open file: %w", err))
	}
	defer func() {
		if err := f.Close(); err != nil {
			panic(fmt.Errorf("failed to close file: %w", err))
		}
	}()

	// Read the file content (example)
	content, err := io.ReadAll(f)
	if err != nil {
		panic(fmt.Errorf("failed to read file: %v", err))
	}

	return content
}
