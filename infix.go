package main

type infix struct{}

// Calculate implements the Calculator interface for expressions given in infix notation.
func (i *infix) Calculate(expr string) (float64, error) {
	expr, err := ConvertInfixToPostfix(expr)
	if err != nil {
		return 0, err
	}

	postfix, err := NewCalculator(PostfixNotation)
	if err != nil {
		return 0, err
	}

	return postfix.Calculate(expr)
}
