package calculator

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

	for _, token := range tokenizePostfixExpression(expr) {
		if len(token) == 0 {
			continue
		}

		op, ok := Operator(token)
		if !ok {
			// attempt to parse float from token and push its value onto stack
			val, err := strconv.ParseFloat(token, 0)
			if err != nil {
				return 0, ErrInvalidPostfixToken
			}
			stack = append(stack, val)
			continue
		}

		top := len(stack)
		if top < 2 {
			return 0, ErrInvalidStackOrdering
		}

		lhs, rhs := stack[top-2], stack[top-1]
		stack = stack[:top-1]
		stack[top-2] = op(lhs, rhs)
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
