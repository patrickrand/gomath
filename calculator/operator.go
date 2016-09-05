package calculator

// Operator is a factory function that maps binary operators to their respective function implementation, as well as returning a boolean indicating whether the mapping exists.
// If the requested mapping does not exist, a 'noop' function that always returns '0' is returned (in addition to 'false').
func Operator(op byte) (func(a, b float64) float64, bool) {
	switch op {
	case '+':
		return func(a, b float64) float64 { return a + b }, true // addition
	case '-':
		return func(a, b float64) float64 { return a - b }, true // subtraction
	case '*':
		return func(a, b float64) float64 { return a * b }, true // multiplication
	case '/':
		return func(a, b float64) float64 { return a / b }, true // division
	default:
		return func(a, b float64) float64 { return 0 }, false // noop
	}
}
