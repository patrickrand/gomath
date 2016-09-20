package main

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"text/scanner"
)

var (
	ErrMissingLeftParenths = errors.New("missing left parenths in expression stack")
	ErrExtraParenths       = errors.New("extra parenths token in expression stack")
)

func push(stack []string, s string) []string {
	stack = append(stack, s)
	fmt.Printf("push %s\t\t= %v\n", s, stack)
	return stack
}

func peek(stack []string) string {
	if len(stack) == 0 {
		return ""
	}
	return stack[0]
}
func pop(stack []string) (string, []string) {
	n := len(stack)
	if n == 0 {
		return "", stack
	}
	top := stack[n-1]
	stack = stack[:n-1]
	fmt.Printf("pop %s\t\t= %v\n", top, stack)
	return top, stack
}

func enqueue(sb *bytes.Buffer, elem string) {
	sb.WriteString(elem)
	fmt.Printf("enqueue %s\t= [%s]\n", elem, sb.String())
	sb.WriteByte(' ')
}

// ConvertInfixToPostfix returns the postfix equivalent of the given infix expression.
func ConvertInfixToPostfix(infix string) (postfix string, err error) {
	fmt.Println("ConvertInfixToPostfix")
	scan := new(scanner.Scanner).Init(strings.NewReader(infix))
	scan.Mode = scanner.ScanIdents | scanner.ScanFloats

	sb := new(bytes.Buffer)
	defer func() {
		postfix = strings.TrimSpace(sb.String())
	}()

	var stack []string

	for tok := scan.Scan(); tok != scanner.EOF; tok = scan.Scan() {
		token := scan.TokenText()
		tokenType := GetTokenType(token)

		switch tokenType {
		case NumberToken:
			// if token == "-" {
			// 	if tok = scan.Scan(); tok == scanner.EOF {
			// 		return postfix, ErrInvalidStackOrdering
			// 	}

			// 	next := scan.TokenText()
			// 	if _, err := strconv.ParseFloat(next, 0); err != nil {
			// 		return postfix, ErrInvalidStackOrdering
			// 	}

			// 	if len(stack) == 0 || !isNumber(prev) {
			// 		sb.WriteString(token)
			// 		sb.WriteString(next)
			// 	} else {
			// 		sb.WriteString(next)
			// 		sb.WriteByte(' ')
			// 		sb.WriteString(token)
			// 	}
			// 	sb.WriteByte(' ')

			// 	token = next
			// 	break
			// }
			enqueue(sb, token)
		case OperatorToken:
			if len(stack) == 0 {
				stack = push(stack, token)
				break
			}

			top := stack[0]
			p1, p2 := precedence(token), precedence(top)

			if (p1 > p2) || (p1 == p2 && isRightAssociativeOperator(token)) {
				stack = push(stack, token)
				break
			}

			//	var enqueued bool
			for (p1 < p2) || (p1 == p2) && isLeftAssociativeOperator(token) {
				enqueue(sb, top)
				//enqueued = true

				top, stack = pop(stack)
				if len(stack) == 0 {
					break
				}
				p2 = precedence(top)
			}
			// if !enqueued {
			stack = push(stack, token)
			// }
		case FunctionToken, LeftParenthsToken:
			stack = push(stack, token)
		case CommaToken, RightParenthsToken:
			var i int
			for i = len(stack) - 1; i >= 0; i-- {
				elem := stack[i]
				if elem == "(" {
					break
				}
				enqueue(sb, elem)
			}

			if i == -1 {
				return postfix, ErrMissingLeftParenths
			}

			if tokenType == CommaToken {
				for j := 0; j < i+1; j++ {
					pop(stack)
				}
				break
			}

			stack = stack[:i]
			elem := peek(stack)
			if _, ok := Function(elem); ok {
				_, stack = pop(stack)
				enqueue(sb, elem)
			}
			// if len(stack) != 0 {
			// 	elem := stack[len(stack)-1]
			// 	if _, ok := Function(elem); ok {
			// 		stack = stack[:len(stack)-1]
			// 		enqueue(sb, elem)
			// 	}
			// }
		default:
			return postfix, ErrInvalidPostfixToken
		}
	}

	var top string
	for len(stack) != 0 {
		top, stack = pop(stack)
		if top == "(" || top == ")" {
			return postfix, ErrExtraParenths
		}
		enqueue(sb, top)
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
