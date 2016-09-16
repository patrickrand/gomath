package main

import "testing"

func TestCalculateInfix(t *testing.T) {
	tests := []struct {
		expr   string
		result float64
		error
	}{
		{expr: "(3+2) + 100 /4", result: 30},
		// {expr: "(3+2) - 100 /4", result: 30},
	}

	i := &infix{}

	for _, tc := range tests {
		if result, err := i.Calculate(tc.expr); result != tc.result || err != tc.error {
			t.Errorf("Calculate(%q) -> { %v, %v} != { %v, %v }", tc.expr, result, err, tc.result, tc.error)
		}
	}

}
