package main

import (
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
	scan.Mode = scanner.ScanIdents | scanner.ScanInts | scanner.ScanFloats

	var stack []string
	var queue []byte
	defer func() { postfix = strings.TrimSpace(string(queue)) }()

	for tok := scan.Scan(); tok != scanner.EOF; tok = scan.Scan() {
		token := scan.TokenText()

		if _, err := strconv.ParseFloat(token, 0); err == nil {
			queue = append(queue, []byte(token+" ")...)
			continue
		}

		if _, ok := Operator(token); ok {
			if len(stack) != 0 {
				o1, o2 := token, stack[len(stack)-1]
				if (isLeftAssociativeOperator(o1) && precedence(o1) <= precedence(o2)) ||
					(isRightAssociativeOperator(o1) && precedence(o1) < precedence(o2)) {
					stack = stack[:len(stack)-1]
					queue = append(queue, []byte(o2+" ")...)
				}
			}

			stack = append(stack, token)
			continue
		}

		if _, ok := Function(token); ok {
			stack = append(stack, token)
			continue
		}

		if token == "," {
			var i int
			for i = len(stack) - 1; i >= 0; i-- {
				elem := stack[i]
				if elem == "(" {
					break
				}
				queue = append(queue, []byte(elem+" ")...)
			}

			if i == -1 {
				return postfix, ErrMissingLeftParenths
			}

			stack = stack[:i+1]
			continue
		}

		if token == "(" {
			stack = append([]string{token}, stack...)
			continue
		}

		if token == ")" {
			var i int
			for i = len(stack) - 1; i >= 0; i-- {
				elem := stack[i]
				if elem == "(" {
					break
				}
				queue = append(queue, []byte(elem+" ")...)
			}

			if i == -1 {
				return postfix, ErrMissingLeftParenths
			}

			stack = stack[:i]

			if len(stack) != 0 {
				elem := stack[len(stack)-1]
				if _, ok := Function(elem); ok {
					stack = stack[:len(stack)-1]
					queue = append(queue, []byte(elem+" ")...)
				}
			}

			continue
		}
	}

	for len(stack) != 0 {
		elem := stack[len(stack)-1]
		if elem == "(" || elem == ")" {
			return postfix, ErrExtraParenths
		}
		stack = stack[:len(stack)-1]
		queue = append(queue, []byte(elem+" ")...)
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
