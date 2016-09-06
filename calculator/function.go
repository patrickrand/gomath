package calculator

import (
	"math"
	"strings"
)

// Function is a factory function that maps keys to their respective function implementation, as well as returning a boolean indicating whether the mapping exists.
// If the requested mapping does not exist, a 'noop' function that always returns '0' is returned (in addition to 'false').
func Function(fn string) (func(x float64) float64, bool) {
	switch strings.ToLower(fn) {
	case "abs":
		return math.Abs, true // absolute-value
	case "sqrt":
		return math.Sqrt, true
	case "cos":
		return math.Cos, true
	case "sin":
		return math.Sin, true
	case "tan":
		return math.Tan, true
	default:
		return func(x float64) float64 { return 0 }, false // noop
	}
}
