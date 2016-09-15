package main

import "testing"

func TestConvertInfixToPostfix(t *testing.T) {
	cases := []struct {
		infix, postfix string
		error
	}{
		{"3 +4 *sin(12)", "3 4 12 sin * +", nil},
	}

	for _, tc := range cases {
		got, err := ConvertInfixToPostfix(tc.infix)
		if tc.error != err {
			t.Errorf("given: %q, expected error: %v, got error: %v", tc.infix, tc.error, err)
		}

		if tc.postfix != got {
			t.Errorf("given: %q, expected: %q, got: %q", tc.infix, tc.postfix, got)
		}
	}
}
