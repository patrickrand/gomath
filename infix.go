package main

import "errors"

type infix struct{}

var (
	ErrInvalidInfixToken = errors.New("invalid infix token")
)

// Calculate implements the Calculator interface for expressions given in infix notation.
func (i *infix) Calculate(expr string) (float64, error) {

	return 0, nil
}
