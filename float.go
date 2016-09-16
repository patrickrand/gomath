package main

import (
	"math"
	"strconv"
)

// Epsilon represents the floating-point precision of any float64 result returned by this package.
const Epsilon float64 = 1e-8

// ParseFloat attempts to parse the given string into a float64 value.
func ParseFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 0)
}

// Equals determines whether the given float64 values are equal up to the precision defined by EPSILON, or share the same (signed) infinity value.
func Equals(a, b float64) bool {
	return BothNaN(a, b) || SameInfinity(a, b) || ((a-b) < Epsilon && (b-a) < Epsilon)
}

// BothNaN returns whether the given float64 values are both equal to NaN.
func BothNaN(a, b float64) bool {
	return math.IsNaN(a) && math.IsNaN(b)
}

// SameInfinity returns whether the given float64 values are equal to the same (signed) infinity value.
func SameInfinity(a, b float64) bool {
	return (math.IsInf(a, 1) && math.IsInf(b, 1)) || (math.IsInf(a, -1) && math.IsInf(b, -1))
}
