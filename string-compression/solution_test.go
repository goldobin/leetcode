package string_compression

import (
	"fmt"
	"math"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_compress(t *testing.T) {
	tests := []struct {
		name      string
		chars     []byte
		wantChars []byte
	}{
		{
			name:      "case 0.1 empty case",
			chars:     []byte(""),
			wantChars: []byte(""),
		},
		{
			name:      "case 0.2 nil case",
			chars:     nil,
			wantChars: nil,
		},
		{
			name:      "case 1.1 leetcode",
			chars:     []byte("aabbccc"),
			wantChars: []byte("a2b2c3"),
		},
		{
			name:      "case 1.2 leetcode",
			chars:     []byte("a"),
			wantChars: []byte("a"),
		},
		{
			name:      "case 1.3 leetcode",
			chars:     []byte("abbbbbbbbbbbb"),
			wantChars: []byte("ab12"),
		},
		{
			name:      "case 2.1",
			chars:     []byte("abcd"),
			wantChars: []byte("abcd"),
		},
		{
			name:      "case 2.2",
			chars:     []byte("bbbbbbbbbbbba"),
			wantChars: []byte("b12a"),
		},
		{
			name:      "case 2.3",
			chars:     []byte("abbbbbbbbbbbba"),
			wantChars: []byte("ab12a"),
		},
		{
			name:      "case 2.4",
			chars:     []byte("bbbbbbbbba"),
			wantChars: []byte("b9a"),
		},
		{
			name:      "case 2.5",
			chars:     []byte("bbbbbbbbbba"),
			wantChars: []byte("b10a"),
		},
		{
			name:      "case 2.6",
			chars:     []byte("abbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbba"),
			wantChars: []byte("ab96a"),
		},
		{
			name:      "case 2.7",
			chars:     []byte("abbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbba"),
			wantChars: []byte("ab99a"),
		},
		{
			name:      "case 2.8",
			chars:     []byte("abbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbba"),
			wantChars: []byte("ab100a"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := slices.Clone(tt.chars)
			gotLen := compress(in)

			assert.Equal(t, len(tt.wantChars), gotLen)

			if gotLen == 0 {
				return
			}

			gotChars := in[:gotLen]
			assert.Equal(t, string(tt.wantChars), string(gotChars))
		})
	}
}

func Test_log10i(t *testing.T) {
	tests := []struct {
		in   int
		want int
	}{
		{in: 0, want: 0},
		{in: 1, want: 0},
		{in: 5, want: 0},
		{in: 7, want: 0},
		{in: 9, want: 0},
		{in: 10, want: 1},
		{in: 11, want: 1},
		{in: 19, want: 1},
		{in: 20, want: 1},
		{in: 99, want: 1},
		{in: 100, want: 2},
		{in: 101, want: 2},
		{in: 999, want: 2},
		{in: 1000, want: 3},
	}

	for _, tt := range tests {
		name := fmt.Sprint(tt.in)
		t.Run(name, func(t *testing.T) {
			got := log10i(tt.in)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Benchmark_log10i(b *testing.B) {
	var (
		n = 0
		v int
	)
	for b.Loop() {
		v = log10i(n % 10000)
		n++
	}
	b.Log(v)
}

func Benchmark_log10(b *testing.B) {
	var (
		n = 0
		v int
	)
	for b.Loop() {
		v = int(math.Log10(float64(n % 10000)))
		n++
	}
	b.Log(v)
}

func Benchmark_compress(b *testing.B) {
	b.ReportAllocs()
	chars := []byte("aabbcccaabbcccaabbcccaabbcccaabbcccaabbcccaabbcccaabbcccaabbcccaabbcccaabbcccaabbcccaabbcccaabbccc")
	var l int
	for b.Loop() {
		l = compress(chars)
	}
	b.Log(l)
}
