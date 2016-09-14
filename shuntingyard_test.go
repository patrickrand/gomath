package main

import "testing"

func TestConvertInfixToPostfix(t *testing.T) {
	cases := []struct {
		infix, postfix string
		error
	}{
		{"3 + 4", "3 4 +", nil},
		{"( 3 + 4 ) * 12 - 5", "3 4 + 12 * 5 -", nil},
	}

	for _, tc := range cases {
		got, err := ConvertPostfixToInfix(tc.infix)
		if tc.error != err {
			t.Errorf("given: %q, expected error: %v, got error: %v", tc.infix, tc.error, err)
		}

		if tc.postfix != got {
			t.Errorf("given: %q, expected: %q, got: %q", tc.infix, tc.postfix, got)
		}
	}
}
