package main

import (
	"math"
	"strings"
)

// Function is a factory function that maps tokens to their respective function implementation, as well as returning a boolean indicating whether the mapping exists.
// If the requested token is invalid, a 'noop' function that always returns '0' is returned (in addition to 'false') to the caller.
func Function(token string) (func(x float64) float64, bool) {
	switch strings.ToLower(token) {
	case "abs":
		return math.Abs, true // absolute value
	case "sqrt":
		return math.Sqrt, true // square root
	case "cos":
		return math.Cos, true // cosine
	case "sin":
		return math.Sin, true // sine
	case "tan":
		return math.Tan, true // tangent
	case "lg":
		return math.Log10, true // common logarithm
	case "lb":
		return math.Log2, true // binary logarithm
	case "ln":
		return math.Log, true // natural logarithm
	default:
		return func(x float64) float64 { return 0 }, false // noop
	}
}
