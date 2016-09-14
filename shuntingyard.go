package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ConvertPostfixToInfix returns the postfix equivalent of the given infix expression.
// Its time-complexity is linear with regards to the length of the given infix string.
func ConvertPostfixToInfix(infix string) (postfix string, err error) {
	var queue []byte
	var stack []string

	defer func() {
		postfix = strings.TrimSpace(string(queue))
	}()

	// TODO
	// scan := new(scanner.Scanner)
	// scan = scan.Init(strings.NewReader(s))
	// scan.Mode = scanner.ScanIdents | scanner.ScanFloats | scanner.ScanInts

	tokens := strings.Split(infix, " ")
	for _, tok := range tokens {
		fmt.Printf("tok: %s\n", tok)
		if isNumber(tok) {
			queue = append(queue, []byte(tok+" ")...)
			continue
		}

		if isOperator(tok) {
			if len(stack) != 0 {
				o1, o2 := tok, stack[0]
				if (isLeftAssociativeOperator(o1) && precedence(o1) <= precedence(o2)) ||
					(isRightAssociativeOperator(o1) && precedence(o1) < precedence(o2)) {
					stack = stack[1:]
					queue = append(queue, []byte(o2+" ")...)
				}
			}

			stack = append([]string{tok}, stack...)
			fmt.Printf("%s -> %#v\n", tok, stack)
			continue
		}

		if isFunction(tok) {
			stack = append([]string{tok}, stack...)
			fmt.Printf("%s -> %#v\n", tok, stack)
			continue
		}

		if isComma(tok) {
			var i int
			for i = 0; i < len(stack); i++ {
				elem := stack[i]
				if isLeftParenths(elem) {
					break
				}
				queue = append(queue, []byte(elem+" ")...)
			}

			if i == len(stack) {
				return postfix, fmt.Errorf("expected left parenths in stack: %#v", stack)
			}

			stack = stack[i:]
			continue
		}

		if isLeftParenths(tok) {
			stack = append([]string{tok}, stack...)
			fmt.Printf("%s -> %#v\n", tok, stack)

			continue
		}

		if isRightParenths(tok) {
			var i int
			for i = 0; i < len(stack); i++ {
				elem := stack[i]
				if isLeftParenths(elem) {
					break
				}
				queue = append(queue, []byte(elem+" ")...)
			}

			if i == len(stack) {
				return postfix, fmt.Errorf("expected left parenths in stack %#v", stack)
			}

			stack = stack[i+1:]

			if len(stack) != 0 {
				if elem := stack[0]; isFunction(elem) {
					stack = stack[1:]
					queue = append(queue, []byte(elem+" ")...)
				}
			}

			continue
		}

	}

	for len(stack) != 0 {
		elem := stack[0]
		if isLeftParenths(elem) || isRightParenths(elem) {
			return postfix, fmt.Errorf("unexpected parenths token remaining on stack %#v", stack)
		}
		stack = stack[1:]
		queue = append(queue, []byte(elem+" ")...)
	}

	return postfix, nil
}

func isNumber(s string) bool {
	_, err := strconv.ParseFloat(s, 0)
	return err == nil
}

func isComma(s string) bool {
	return s == ","
}

func isOperator(s string) bool {
	_, ok := Operator(s)
	return ok
}

func isFunction(s string) bool {
	_, ok := Function(s)
	return ok
}

func isLeftParenths(s string) bool {
	return s == "("
}

func isRightParenths(s string) bool {
	return s == ")"
}

func isLeftAssociativeOperator(s string) bool {
	return s != "pow" && isOperator(s)
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
