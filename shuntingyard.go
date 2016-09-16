package main

import (
	"bytes"
	"errors"
	"strconv"
	"strings"
	"text/scanner"
)

var (
	ErrMissingLeftParenths = errors.New("missing left parenths in expression stack")
	ErrExtraParenths       = errors.New("extra parenths token in expression stack")
)

type TokenType int

const (
	NumberToken TokenType = iota
	OperatorToken
	FunctionToken
	CommaToken
	LeftParenthsToken
	RightParenthsToken
	ErrorToken
)

func GetTokenType(token string) TokenType {
	if _, err := strconv.ParseFloat(token, 0); err == nil {
		return NumberToken
	}

	if _, ok := Operator(token); ok {
		return OperatorToken
	}

	if _, ok := Function(token); ok {
		return FunctionToken
	}

	if token == "," {
		return CommaToken
	}

	if token == "(" {
		return LeftParenthsToken
	}

	if token == ")" {
		return RightParenthsToken
	}

	return ErrorToken
}

// ConvertInfixToPostfix returns the postfix equivalent of the given infix expression.
func ConvertInfixToPostfix(infix string) (postfix string, err error) {
	scan := new(scanner.Scanner).Init(strings.NewReader(infix))
	scan.Mode = scanner.ScanIdents | scanner.ScanFloats

	sb := new(bytes.Buffer)
	defer func() {
		postfix = strings.TrimSpace(sb.String())
	}()

	var stack []string
	var prev string

	for tok := scan.Scan(); tok != scanner.EOF; tok = scan.Scan() {
		token := scan.TokenText()

		switch GetTokenType(token) {
		case NumberToken:
			sb.WriteString(token)
			sb.WriteByte(' ')
		case OperatorToken:
			if len(stack) != 0 {
				if token == "-" {
					next := scan.Scan()

					text := scan.TokenText()

					if next == scanner.EOF {
						return postfix, ErrInvalidStackOrdering
					}

					if _, err := strconv.ParseFloat(text, 0); err != nil {
						return postfix, ErrInvalidStackOrdering
					}

					if isNumber(prev) {
						sb.WriteString(text)
						sb.WriteByte(' ')
						sb.WriteString(token)
					} else {
						sb.WriteString(token)
						sb.WriteString(text)
					}
					prev = token
					tok = next
					token = text

					sb.WriteByte(' ')
					continue
				}
			} else if token == "-" {
				tok = scan.Scan()
				if tok == scanner.EOF {
					return postfix, ErrInvalidStackOrdering
				}

				if _, err := strconv.ParseFloat(scan.TokenText(), 0); err != nil {
					return postfix, ErrInvalidStackOrdering
				}

				sb.WriteString(token)
				sb.WriteString(scan.TokenText())
				sb.WriteByte(' ')
				continue
			}
			stack = append(stack, token)
		case FunctionToken:
			stack = append(stack, token)
		case CommaToken:
			var i int
			for i = len(stack) - 1; i >= 0; i-- {
				elem := stack[i]
				if elem == "(" {
					break
				}
				sb.WriteString(elem)
				sb.WriteByte(' ')
			}

			if i == -1 {
				return postfix, ErrMissingLeftParenths
			}

			stack = stack[:i+1]
		case LeftParenthsToken:
			stack = append([]string{token}, stack...)
		case RightParenthsToken:
			var i int
			for i = len(stack) - 1; i >= 0; i-- {
				elem := stack[i]
				if elem == "(" {
					break
				}

				sb.WriteString(elem)
				sb.WriteByte(' ')
			}

			if i == -1 {
				return postfix, ErrMissingLeftParenths
			}

			stack = stack[:i]

			if len(stack) != 0 {
				elem := stack[len(stack)-1]
				if _, ok := Function(elem); ok {
					stack = stack[:len(stack)-1]
					sb.WriteString(elem)
					sb.WriteByte(' ')
				}
			}
		default:
			return postfix, ErrInvalidPostfixToken
		}
		prev = token
	}

	for len(stack) != 0 {
		elem := stack[len(stack)-1]
		if elem == "(" || elem == ")" {
			return postfix, ErrExtraParenths
		}
		stack = stack[:len(stack)-1]
		sb.WriteString(elem)
		sb.WriteByte(' ')
	}

	return postfix, nil
}

func isNumber(s string) bool {
	_, err := strconv.ParseFloat(s, 0)
	return err == nil
}

func isLeftAssociativeOperator(s string) bool {
	_, ok := Operator(s)
	return ok && s != "pow"
}

func isRightAssociativeOperator(s string) bool {
	return s == "pow"
}

func precedence(op string) int {
	switch op {
	case "+", "-":
		return 0
	case "*", "/":
		return 1
	case "pow", "root":
		return 2
	default:
		return -1
	}
}
