package main

import (
	"strconv"
	"strings"
)

func shuntingYard(infix string) (postfix string, err error) {
	var output []byte

	defer func() {
		postfix = string(output)
	}()

	tokens := strings.Split(infix, " ")
	for _, tok := range tokens {
		if isNumber(tok) {
			output = append(output, []byte(tok+" ")...)
			continue
		}

		if isOperator(tok) {
			// stuff
			continue
		}

		if isFunction(tok) {
			// stuff
			continue
		}

		if isComma(tok) {
			// stuff
			continue
		}

		if isLeftParenths(tok) {
			continue
		}

		if isRightParenths(tok) {
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
