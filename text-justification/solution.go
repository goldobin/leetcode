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
	)

	for {
		if wordIx >= wordsTotal {
			line := joinAndRightPad(words[firstWordInLineIx:], maxWidth)
			result = append(result, line)
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

		line := justifyLine(words[firstWordInLineIx:wordIx], maxWidth)
		result = append(result, line)
		lineLen = 0
	}

	return result
}

func joinAndRightPad(words []string, maxWidth int) string {
	if len(words) == 0 {
		panic("empty words")
	}

	line := strings.Join(words, " ")
	if len(line) > maxWidth {
		panic(fmt.Sprintf("words do not fit into max with %d", maxWidth))
	}

	spacesToAdd := maxWidth - len(line)
	if spacesToAdd <= 0 {
		return line
	}

	spaces := strings.Repeat(" ", spacesToAdd)
	return line + spaces
}

func justifyLine(words []string, maxWidth int) string {
	if len(words) == 0 {
		panic("empty words")
	}

	sb := strings.Builder{}
	sb.Grow(maxWidth)

	if len(words) == 1 {
		word := words[0]
		if len(word) > maxWidth {
			panic(fmt.Sprintf("word len %d exceeds max width %d", len(word), maxWidth))
		}

		sb.WriteString(word)
		spaceCount := maxWidth - len(word)
		for i := 0; i < spaceCount; i++ {
			sb.WriteByte(' ')
		}

		return sb.String()
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

	return sb.String()
}
