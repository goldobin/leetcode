package text_justification

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fullJustify(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		maxWidth int
		want     []string
	}{
		{
			name:     "case 0.1.1 nil",
			words:    nil,
			maxWidth: 1,
			want:     nil,
		},
		{
			name:     "case 0.1.1 empty",
			words:    []string{},
			maxWidth: 1,
			want:     nil,
		},
		{
			name:     "case 0.2.1 one word",
			words:    []string{"word"},
			maxWidth: 4,
			want:     []string{"word"},
		},
		{
			name:     "case 0.2.2 one word",
			words:    []string{"word"},
			maxWidth: 5,
			want:     []string{"word "},
		},
		{
			name:     "case 0.2.3 one word",
			words:    []string{"word"},
			maxWidth: 8,
			want:     []string{"word    "},
		},
		{
			name:     "case 0.4 one word per line",
			words:    []string{"word", "word2"},
			maxWidth: 5,
			want: []string{
				"word ",
				"word2",
			},
		},
		{
			name:     "case 0.5.1 two words",
			words:    []string{"some", "word"},
			maxWidth: 9,
			want: []string{
				"some word",
			},
		},
		{
			name:     "case 0.5.2 two words",
			words:    []string{"some", "word"},
			maxWidth: 10,
			want: []string{
				"some word ",
			},
		},
		{
			name:     "case 0.5.3 two words",
			words:    []string{"some", "word"},
			maxWidth: 11,
			want: []string{
				"some word  ",
			},
		},
		{
			name:     "case 0.5.4 two words",
			words:    []string{"some", "word"},
			maxWidth: 12,
			want: []string{
				"some word   ",
			},
		},
		{
			name:     "case 0.6 two words per line",
			words:    []string{"some", "word", "other", "wrd"},
			maxWidth: 9,
			want: []string{
				"some word",
				"other wrd",
			},
		},
		{
			name:     "case 0.7 two words per line and one",
			words:    []string{"some", "word", "other"},
			maxWidth: 9,
			want: []string{
				"some word",
				"other    ",
			},
		},
		{
			name:     "case 0.8 two words per",
			words:    []string{"some", "word", "and", "one"},
			maxWidth: 9,
			want: []string{
				"some word",
				"and one  ",
			},
		},
		{
			name:     "case 1.1 leetcode",
			words:    []string{"This", "is", "an", "example", "of", "text", "justification."},
			maxWidth: 16,
			want: []string{
				"This    is    an",
				"example  of text",
				"justification.  ",
			},
		},
		{
			name:     "case 1.2 leetcode",
			words:    []string{"What", "must", "be", "acknowledgment", "shall", "be"},
			maxWidth: 16,
			want: []string{
				"What   must   be",
				"acknowledgment  ",
				"shall be        ",
			},
		},
		{
			name: "case 1.3 leetcode",
			words: []string{
				"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.",
				"Art", "is", "everything", "else", "we", "do",
			},
			maxWidth: 20,
			want: []string{
				"Science  is  what we",
				"understand      well",
				"enough to explain to",
				"a  computer.  Art is",
				"everything  else  we",
				"do                  ",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fullJustify(tt.words, tt.maxWidth)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_joinAndRightPadWhenWordsDoNotFit(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		maxWidth int
		want     string
	}{
		{
			name:     "case 0.1 empty",
			words:    []string{},
			maxWidth: 1,
			want:     "",
		},
		{
			name:     "case 1.1 one word",
			words:    []string{"word"},
			maxWidth: 1,
		},
		{
			name:     "case 1.2 one word",
			words:    []string{"word"},
			maxWidth: 3,
		},
		{
			name:     "case 2.1 two words",
			words:    []string{"what", "to"},
			maxWidth: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Panics(t, func() {
				joinAndRightPad(&strings.Builder{}, tt.words, tt.maxWidth)
			})
		})
	}
}

func Test_joinAndRightPad(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		maxWidth int
		want     string
	}{
		{
			name:     "case 1.1 one word",
			words:    []string{"word"},
			maxWidth: 4,
			want:     "word",
		},
		{
			name:     "case 1.2 one word",
			words:    []string{"word"},
			maxWidth: 5,
			want:     "word ",
		},
		{
			name:     "case 1.3 one word",
			words:    []string{"word"},
			maxWidth: 7,
			want:     "word   ",
		},
		{
			name:     "case 2.1 two words",
			words:    []string{"word", "do"},
			maxWidth: 7,
			want:     "word do",
		},
		{
			name:     "case 2.2 two words",
			words:    []string{"word", "do"},
			maxWidth: 8,
			want:     "word do ",
		},
		{
			name:     "case 2.2 two words",
			words:    []string{"word", "do"},
			maxWidth: 11,
			want:     "word do    ",
		},
		{
			name:     "case 3.1 three words",
			words:    []string{"what", "to", "do"},
			maxWidth: 10,
			want:     "what to do",
		},
		{
			name:     "case 3.2 three words",
			words:    []string{"what", "to", "do"},
			maxWidth: 11,
			want:     "what to do ",
		},
		{
			name:     "case 3.3 three words",
			words:    []string{"what", "to", "do"},
			maxWidth: 12,
			want:     "what to do  ",
		},
	}

	for _, tt := range tests {
		sb := strings.Builder{}
		sb.Grow(tt.maxWidth)
		t.Run(tt.name, func(t *testing.T) {
			sb.Reset()
			joinAndRightPad(&sb, tt.words, tt.maxWidth)
			got := sb.String()

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_justifyLineWhenWordsDoNotFit(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		maxWidth int
		want     string
	}{
		{
			name:     "case 0.1 empty",
			words:    []string{},
			maxWidth: 1,
			want:     "",
		},
		{
			name:     "case 1.1 one word",
			words:    []string{"word"},
			maxWidth: 1,
		},
		{
			name:     "case 1.2 one word",
			words:    []string{"word"},
			maxWidth: 3,
		},
		{
			name:     "case 2.1 two words",
			words:    []string{"what", "to"},
			maxWidth: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Panics(t, func() {
				joinAndRightPad(&strings.Builder{}, tt.words, tt.maxWidth)
			})
		})
	}
}

func Test_justifyLine(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		maxWidth int
		want     string
	}{
		{
			name:     "case 1.1 one word",
			words:    []string{"word"},
			maxWidth: 4,
			want:     "word",
		},
		{
			name:     "case 1.2 one word",
			words:    []string{"word"},
			maxWidth: 5,
			want:     "word ",
		},
		{
			name:     "case 1.3 one word",
			words:    []string{"word"},
			maxWidth: 6,
			want:     "word  ",
		},
		{
			name:     "case 2.1 two words",
			words:    []string{"what", "the"},
			maxWidth: 8,
			want:     "what the",
		},
		{
			name:     "case 2.2 two words",
			words:    []string{"what", "the"},
			maxWidth: 9,
			want:     "what  the",
		},
		{
			name:     "case 2.2 two words",
			words:    []string{"what", "the"},
			maxWidth: 11,
			want:     "what    the",
		},
		{
			name:     "case 3.1 three words",
			words:    []string{"what", "the", "hell"},
			maxWidth: 13,
			want:     "what the hell",
		},
		{
			name:     "case 3.2 three words",
			words:    []string{"what", "the", "hell"},
			maxWidth: 14,
			want:     "what  the hell",
		},
		{
			name:     "case 3.3 three words",
			words:    []string{"what", "the", "hell"},
			maxWidth: 15,
			want:     "what  the  hell",
		},
		{
			name:     "case 3.4 three words",
			words:    []string{"what", "the", "hell"},
			maxWidth: 16,
			want:     "what   the  hell",
		},
		{
			name:     "case 3.4 three words",
			words:    []string{"what", "the", "hell"},
			maxWidth: 17,
			want:     "what   the   hell",
		},
		{
			name:     "case 4.1 four words",
			words:    []string{"what", "the", "hell", "is"},
			maxWidth: 16,
			want:     "what the hell is",
		},
		{
			name:     "case 4.2 four words",
			words:    []string{"what", "the", "hell", "is"},
			maxWidth: 17,
			want:     "what  the hell is",
		},
		{
			name:     "case 4.3 four words",
			words:    []string{"what", "the", "hell", "is"},
			maxWidth: 18,
			want:     "what  the  hell is",
		},
		{
			name:     "case 4.4 four words",
			words:    []string{"what", "the", "hell", "is"},
			maxWidth: 19,
			want:     "what  the  hell  is",
		},
		{
			name:     "case 4.5 four words",
			words:    []string{"what", "the", "hell", "is"},
			maxWidth: 20,
			want:     "what   the  hell  is",
		},
		{
			name:     "case 4.6 four words",
			words:    []string{"what", "the", "hell", "is"},
			maxWidth: 21,
			want:     "what   the   hell  is",
		},
		{
			name:     "case 4.7 four words",
			words:    []string{"what", "the", "hell", "is"},
			maxWidth: 22,
			want:     "what   the   hell   is",
		},
		{
			name:     "case 4.8 four words",
			words:    []string{"what", "the", "hell", "is"},
			maxWidth: 23,
			want:     "what    the   hell   is",
		},
	}

	for _, tt := range tests {
		sb := strings.Builder{}
		sb.Grow(tt.maxWidth)
		t.Run(tt.name, func(t *testing.T) {
			sb.Reset()
			justifyLine(&sb, tt.words, tt.maxWidth)
			got := sb.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
