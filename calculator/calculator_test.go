package calculator

import "testing"

func TestNew(t *testing.T) {
	testCases := []struct {
		notation
		Calculator
		error
	}{
		{notation: notation("invalid"), error: ErrNotationNotImplemented},
		{notation: PREFIX, error: ErrNotationNotImplemented},
		{notation: INFIX, error: ErrNotationNotImplemented},
		{notation: POSTFIX, Calculator: &postfix{}},
	}

	for _, tc := range testCases {
		calc, err := New(tc.notation)
		if expected := tc.Calculator; expected == nil && calc != expected {
			t.Errorf("expected nil Calculator for notation %q", string(tc.notation))
		}
		if err != tc.error {
			t.Errorf("expected error to be '%v' for notation %q", err, string(tc.notation))
		}
	}
}
