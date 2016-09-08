package main

import "math"

// Operator is a factory function that maps tokens to their respective function implementation, as well as returning a boolean indicating whether the mapping exists.
// If the requested mapping does not exist, a 'noop' function that always returns '0' is returned (in addition to 'false').
func Operator(token string) (func(a, b float64) float64, bool) {
	switch token {
	case "+":
		return func(a, b float64) float64 { return a + b }, true // addition
	case "-":
		return func(a, b float64) float64 { return a - b }, true // subtraction
	case "*":
		return func(a, b float64) float64 { return a * b }, true // multiplication
	case "/":
		return func(a, b float64) float64 { return a / b }, true // division
	case "pow":
		return math.Pow, true // exponentation
	default:
		return func(a, b float64) float64 { return 0 }, false // noop
	}
}
