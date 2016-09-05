package calculator

import "errors"

type operator func(a, b float64) float64

func OperatorFactory(op byte) (operator, error) {
	switch op {
	case '+':
		return add, nil
	case '-':
		return sub, nil
	case '*':
		return mult, nil
	case '/':
		return div, nil
	default:
		return nil, errors.New("invalid operator: " + string(op))
	}
}

func add(a, b float64) float64 {
	return a + b
}

func sub(a, b float64) float64 {
	return a - b
}

func mult(a, b float64) float64 {
	return a * b
}

func div(a, b float64) float64 {
	return a / b
}
