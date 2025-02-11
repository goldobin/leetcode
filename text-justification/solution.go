package text_justification

import (
	"fmt"
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
	if len(words) == 0 {
		return nil
	}

	var (
		wordsTotal        = len(words)
		result            = make([]string, 0, len(words)/2)
		wordIx            = 0
		firstWordInLineIx = 0
		lineLen           = 0
		sb                = strings.Builder{}
	)

	sb.Grow(maxWidth)

	for {
		if wordIx >= wordsTotal {
			sb.Reset()
			joinAndRightPad(&sb, words[firstWordInLineIx:], maxWidth)
			result = append(result, sb.String())
			break
		}

		word := words[wordIx]
		wordLen := len(word)
		if wordLen > maxWidth {
			panic(fmt.Sprintf("word at position %d exceeds max width %d", wordIx, maxWidth))
		}

		if lineLen == 0 {
			lineLen += wordLen
			firstWordInLineIx = wordIx
			wordIx += 1
			continue
		}

		wl := wordLen + 1
		if lineLen+wl <= maxWidth {
			lineLen += wl
			wordIx += 1
			continue
		}

		sb.Reset()
		justifyLine(&sb, words[firstWordInLineIx:wordIx], maxWidth)
		result = append(result, sb.String())
		lineLen = 0
	}

	return result
}

func joinAndRightPad(sb *strings.Builder, words []string, maxWidth int) {
	if len(words) == 0 {
		panic("empty words")
	}

	lineLen := 0
	for i, word := range words {
		sb.WriteString(word)
		lineLen += len(word)
		if i != len(words)-1 {
			sb.WriteByte(' ')
			lineLen += 1
		}
	}

	if lineLen > maxWidth {
		panic(fmt.Sprintf("words do not fit into max with %d", maxWidth))
	}

	spacesToAdd := maxWidth - lineLen
	if spacesToAdd <= 0 {
		return
	}

	for i := 0; i < spacesToAdd; i++ {
		sb.WriteByte(' ')
	}
}

func justifyLine(sb *strings.Builder, words []string, maxWidth int) {
	if len(words) == 0 {
		panic("empty words")
	}

	if len(words) == 1 {
		word := words[0]
		if len(word) > maxWidth {
			panic(fmt.Sprintf("word len %d exceeds max width %d", len(word), maxWidth))
		}
		rightPadWord(sb, word, maxWidth)
		return
	}

	wordSpaceCount := len(words) - 1
	totalWordLen := 0
	for _, word := range words {
		totalWordLen += len(word)
	}

	totalSpaceLen := maxWidth - totalWordLen
	guaranteedSpaceLen := totalSpaceLen / wordSpaceCount
	additionalSpaceLen := totalSpaceLen - guaranteedSpaceLen*wordSpaceCount

	for i, word := range words {
		sb.WriteString(word)
		if i >= len(words)-1 {
			break
		}

		for j := 0; j < guaranteedSpaceLen; j++ {
			sb.WriteByte(' ')
		}

		if additionalSpaceLen > 0 {
			sb.WriteByte(' ')
			additionalSpaceLen -= 1
		}
	}
}

func rightPadWord(sb *strings.Builder, word string, maxWidth int) {
	sb.WriteString(word)
	spaceCount := maxWidth - len(word)
	for i := 0; i < spaceCount; i++ {
		sb.WriteByte(' ')
	}
}
