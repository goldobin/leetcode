package basic_calculator

import (
	"fmt"
	"strconv"
)

func calculate(s string) int {
	return 0
}

type (
	op        func() int
	tokenType int
	token     struct {
		tt tokenType
		v  int
	}
)

const (
	tokenTypeWhitespace tokenType = iota
	tokenTypePlus
	tokenTypeMinus
	tokenTypeOpenGroup
	tokenTypeCloseGroup
	tokenTypeNumber
)

var (
	whitespaceToken = token{tt: tokenTypeWhitespace}
	plusToken       = token{tt: tokenTypePlus}
	minusToken      = token{tt: tokenTypeMinus}
	openGroupToken  = token{tt: tokenTypeOpenGroup}
	closeGroupToken = token{tt: tokenTypeCloseGroup}
)

func numberToken(n int) token {
	return token{
		tt: tokenTypeNumber,
		v:  n,
	}
}

func parse(s string) ([]token, error) {
	result := make([]token, 0, len(s))

	for {
		offset, t, err := parseNextToken(s)
		if err != nil {
			return nil, err
		}

		if t.tt != tokenTypeWhitespace {
			result = append(result, t)
		}

		if offset >= len(s) {
			break
		}
		s = s[offset:]
	}

	return result, nil
}

func parseNextToken(s string) (int, token, error) {
	if len(s) == 0 {
		return 0, token{}, fmt.Errorf("empty string")
	}

	i := 0
	firstCh := s[0]

	switch firstCh {
	case '+':
		return 1, plusToken, nil
	case '-':
		return 1, minusToken, nil
	case '(':
		return 1, openGroupToken, nil
	case ')':
		return 1, closeGroupToken, nil
	case ' ':
		i++
		for ; i < len(s); i++ {
			if s[i] != ' ' {
				break
			}
		}
		return i, whitespaceToken, nil

	default:
		if firstCh < '0' || firstCh > '9' {
			return 0, token{}, fmt.Errorf("unsupported char %c", firstCh)
		}

		i++
		for ; i < len(s); i++ {
			if s[i] < '0' || s[i] > '9' {
				break
			}
		}

		number := s[:i]
		n, err := strconv.Atoi(number)

		if err != nil {
			return 0, token{}, fmt.Errorf("invalid number %s", number)
		}

		return i, numberToken(n), nil
	}
}

func unaryNegation(o op) op {
	return func() int {
		return -o()
	}
}

func num(n int) op {
	return func() int {
		return n
	}
}
func plus(o1, o2 op) op {
	return func() int {
		return o1() + o2()
	}
}

func minus(o1, o2 op) op {
	return plus(o1, unaryNegation(o2))
}
