package basic_calculator

import (
	"fmt"
	"strconv"
)

func calculate(s string) int {
	tokens, err := parseTokens(s)
	if err != nil {
		panic(err)
	}
	tree, tail, err := parseTree(tokens, false)
	if err != nil {
		panic(fmt.Errorf("failed to parse tree: %w", err))
	}

	if len(tail) > 0 {
		panic(fmt.Sprintf("not all tokens parsed: %v", tail))
	}
	return tree.eval()
}

type (
	tokenType int
	token     struct {
		tt tokenType
		v  int
	}
	node interface {
		fmt.Stringer
		eval() int
	}
	sum      struct{ a, b node }
	negation struct{ a node }
	operand  struct{ v int }
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

func (tt tokenType) String() string {
	switch tt {
	case tokenTypeWhitespace:
		return "whitespace"
	case tokenTypePlus:
		return "plus"
	case tokenTypeMinus:
		return "minus"
	case tokenTypeOpenGroup:
		return "open_group"
	case tokenTypeCloseGroup:
		return "close_group"
	case tokenTypeNumber:
		return "number"
	default:
		return "unknown"
	}
}

func (t token) String() string {
	if t.tt == tokenTypeNumber {
		return fmt.Sprintf("{%v, %v}", t.tt, t.v)
	}

	return t.tt.String()
}

func numberToken(n int) token {
	return token{
		tt: tokenTypeNumber,
		v:  n,
	}
}

func (o sum) eval() int {
	return o.a.eval() + o.b.eval()
}

func (o sum) String() string {
	return fmt.Sprintf("(%v + %v)", o.a, o.b)
}

func (o negation) eval() int {
	return -o.a.eval()
}

func (o negation) String() string {
	return fmt.Sprintf("-%v", o.a)
}

func (o operand) eval() int {
	return o.v
}

func (o operand) String() string {
	return fmt.Sprintf("%v", o.v)
}

func parseOperation(a node, tokens []token) (node, []token, error) {
	if len(tokens) == 0 {
		return nil, tokens, fmt.Errorf("not enogh tokens to parse operation")
	}

	var (
		b      node
		rest   []token
		err    error
		result node
		head   = tokens[0]
		tail   = tokens[1:]
	)

	switch head.tt {
	case tokenTypePlus:
		b, rest, err = parseTree(tail, false)
		if err != nil {
			return nil, tokens, err
		}

		result = sum{a, b}

	case tokenTypeMinus:
		b, rest, err = parseTree(tokens, false)
		if err != nil {
			return nil, tokens, err
		}

		result = sum{a, b}

	default:
		return nil, tokens, fmt.Errorf("unsupported token type for operation '%v'", head.tt)
	}

	return result, rest, nil
}

func parseTree(ts []token, negate bool) (node, []token, error) {
	if len(ts) == 0 {
		return nil, ts, fmt.Errorf("empty token list")
	}

	var (
		head = ts[0]
		tail = ts[1:]
	)

	switch head.tt {
	case tokenTypeMinus:
		return parseTree(tail, true)

	case tokenTypeNumber:
		var (
			a    node = operand{head.v}
			rest      = tail
		)

		if negate {
			a = negation{a}
		}

		if len(rest) > 0 && (rest[0].tt == tokenTypePlus || rest[0].tt == tokenTypeMinus) {
			return parseOperation(a, rest)
		}
		return a, rest, nil

	case tokenTypeOpenGroup:
		a, rest, err := parseTree(tail, false)
		if err != nil {
			return nil, ts, err
		}

		if len(rest) == 0 {
			return nil, ts, fmt.Errorf("empty tail")
		}
		if rest[0].tt != tokenTypeCloseGroup {
			return nil, ts, fmt.Errorf("expected close group")
		}

		if negate {
			a = negation{a}
		}

		rest = rest[1:]
		if len(rest) > 0 && (rest[0].tt == tokenTypePlus || rest[0].tt == tokenTypeMinus) {
			return parseOperation(a, rest)
		}

		return a, rest, nil

	default:
		return nil, ts, fmt.Errorf("unsupported token type '%v'", head.tt)
	}
}

func parseTokens(s string) ([]token, error) {
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

	var (
		i    = 0
		head = s[0]
	)

	switch head {
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
		if head < '0' || head > '9' {
			return 0, token{}, fmt.Errorf("unsupported char '%c'", head)
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
