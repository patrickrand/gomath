package calculator

import "testing"

func TestCalculatePostfix(t *testing.T) {
	tests := []struct {
		expr   string
		result float64
		error
	}{
		{expr: "4 2.3422 + 22.4 * 16 - 3 / .02176 -", result: 42},
		{expr: "4,2.3422   + 22.4,, ,* 16 - 3 /, .02176 -", result: 42},
		{expr: "4 + 9 * 22.4 -", result: 0, error: ErrInvalidStackOrdering},
		{expr: "4 2.3422 + 9 * 22.4", result: 0, error: ErrInvalidStackOrdering},
		{expr: "4 2.3@22 + 9 * 22.4 -", result: 0, error: ErrInvalidPostfixToken},
	}

	pf := &postfix{}

	for _, tc := range tests {
		if result, err := pf.Calculate(tc.expr); result != tc.result || err != tc.error {
			t.Errorf("Calculate(%q) -> { %v, %v} != { %v, %v }", tc.expr, result, err, tc.result, tc.error)
		}
	}

}
