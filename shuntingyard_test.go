package main

import "testing"

func TestConvertInfixToPostfix(t *testing.T) {
	cases := []struct {
		infix, postfix string
		error
	}{
		{"3 +4 *sin(12)", "3 4 12 sin * +", nil},
		{"-3 +4", "-3 4 +", nil},
		{"-3 + -4", "-3 -4 +", nil},
		{"3 + -4 *sin(12)", "3 -4 12 sin * +", nil},
		{"3 * (9 - 5)", "3 9 5 - *", nil},
		{"-2 + (3 * (9 - 5)) + -1", "-2 3 9 5 - * -1 + +", nil},
		// {"-2 -(3 * (9 - 5)) + -1", "-2 3 9 5 - * -1 + +", nil},
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
