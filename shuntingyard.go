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
		tokenType := GetTokenType(token)

		switch tokenType {
		case NumberToken:
			sb.WriteString(token)
			sb.WriteByte(' ')
		case OperatorToken:
			if token != "-" {
				stack = append(stack, token)
				break
			}

			if tok = scan.Scan(); tok == scanner.EOF {
				return postfix, ErrInvalidStackOrdering
			}

			next := scan.TokenText()
			if _, err := strconv.ParseFloat(next, 0); err != nil {
				return postfix, ErrInvalidStackOrdering
			}

			if len(stack) == 0 || !isNumber(prev) {
				sb.WriteString(token)
				sb.WriteString(next)
			} else {
				sb.WriteString(next)
				sb.WriteByte(' ')
				sb.WriteString(token)
			}
			sb.WriteByte(' ')

			token = next
		case FunctionToken, LeftParenthsToken:
			stack = append(stack, token)
		case CommaToken, RightParenthsToken:
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

			if tokenType == CommaToken {
				stack = stack[:i+1]
				break
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
