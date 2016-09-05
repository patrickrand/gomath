package calculator

import (
	"errors"
	"fmt"
	"strconv"
)

type postfix struct {
	raw   string
	stack []string
}

// Calculate implements the Calculator interface for postfix (Reverse Polish) notation.
func (pf *postfix) Calculate() (float64, error) {
	var stack []float64
	for i := 0; i < len(pf.stack); i++ {
		elem := pf.stack[i]

		if len(elem) == 0 {
			continue
		}

		if !isOperator(elem) {
			val, err := strconv.ParseFloat(elem, 0)
			if err != nil {
				return 0, fmt.Errorf("invalid numeric value for postfix evaluation: %s", elem)
			}
			stack = append(stack, val)
			continue
		}

		op, _ := OperatorFactory(elem[0]) // ignore parse error due to preceeding code block

		top := len(stack)
		if top < 2 {
			return 0, fmt.Errorf("invalid postfix order of operations: %v", stack)
		}

		lhs, rhs := stack[top-2], stack[top-1]
		stack = stack[:top-1]
		stack[top-2] = op(lhs, rhs)
	}

	if len(stack) != 1 {
		return 0, errors.New("too many values given for postfix evaluation")
	}

	return stack[0], nil
}

func parsePostfix(expr string) (*postfix, error) {
	var stack []string
	for i := 0; i < len(expr); i++ {
		j := i
		for j < len(expr) {
			if token := expr[j]; token == ',' || token == ' ' {
				break
			}
			j++
		}
		if i != j {
			stack = append(stack, string(expr[i:j]))
		}
		i = j
	}
	return &postfix{raw: expr, stack: stack}, nil
}

func isOperator(s string) bool {
	if len(s) != 1 {
		return false
	}
	switch s[0] {
	case '+', '-', '*', '/':
		return true
	default:
		return false
	}
}
