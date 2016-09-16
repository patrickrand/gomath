package main

import (
	"errors"
	"strconv"
)

type postfix struct{}

var (
	ErrInvalidPostfixToken  = errors.New("invalid postfix token")
	ErrInvalidStackOrdering = errors.New("invalid stack ordering")
)

// Calculate implements the Calculator interface for expressions given in postfix (Reverse Polish) notation.
func (pf *postfix) Calculate(expr string) (float64, error) {
	var stack []float64

	tokens := tokenizePostfixExpression(expr)
	for i, token := range tokens {
		if len(token) == 0 {
			continue
		}

		top := len(stack)

		if op, ok := Operator(token); ok {
			if top <= 1 {
				if top == 1 && token == "-" && isNumber(tokens[i+1]) {
					tokens[i+1] = tokens[i] + tokens[i+1]
					continue
				}
				return 0, ErrInvalidStackOrdering
			}

			// if top == 1 && (token == "+" || token == "-") {
			// 	stack = append(stack, 0)
			// 	top++
			// }

			lhs, rhs := stack[top-2], stack[top-1]
			stack, top = stack[:top-1], top-1
			stack[top-1] = op(lhs, rhs)
			continue
		}

		if fn, ok := Function(token); ok {
			if top < 1 {
				return 0, ErrInvalidStackOrdering
			}

			stack[top-1] = fn(stack[top-1])
			continue
		}

		// attempt to parse float from token and push its value onto stack
		val, err := strconv.ParseFloat(token, 0)
		if err != nil {
			return 0, ErrInvalidPostfixToken
		}
		stack = append(stack, val)
		continue

	}

	if len(stack) != 1 {
		return 0, ErrInvalidStackOrdering
	}

	return stack[0], nil
}

func tokenizePostfixExpression(expr string) []string {
	var stack []string
	for i := 0; i < len(expr); i++ {
		j := i
		for ; j < len(expr); j++ {
			if token := expr[j]; token == ',' || token == ' ' {
				break
			}
		}
		if i != j {
			stack = append(stack, string(expr[i:j]))
		}
		i = j
	}
	return stack
}
