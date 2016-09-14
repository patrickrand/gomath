package main

import (
	"fmt"
	"strconv"
	"strings"
)

func shuntingYard(infix string) (postfix string, err error) {
	var queue []byte
	var stack []string

	defer func() {
		postfix = string(queue)
	}()

	tokens := strings.Split(infix, " ")
	for _, tok := range tokens {
		if isNumber(tok) {
			queue = append(queue, []byte(tok+" ")...)
			continue
		}

		if isOperator(tok) {
			// stuff
			continue
		}

		if isFunction(tok) {
			stack = append([]string{tok}, stack...)
			continue
		}

		if isComma(tok) {
			var i int
			var hasLeftParenths bool
			for i = 0; i < len(stack); i++ {
				elem := stack[i]
				if isLeftParenths(elem) {
					hasLeftParenths = true
					break
				}
				queue = append(queue, []byte(elem+" ")...)
			}
			stack = stack[i:]

			if !hasLeftParenths {
				return postfix, fmt.Errorf("expected left parenths in stack: %#v", stack)
			}
			continue
		}

		if isLeftParenths(tok) {
			stack = append([]string{tok}, stack...)
			continue
		}

		if isRightParenths(tok) {
			// stuff
			continue
		}

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
