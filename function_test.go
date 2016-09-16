package main

import "testing"

func TestFunction(t *testing.T) {
	tests := []struct {
		fn     string
		x      float64
		result float64
		ok     bool
	}{
		{fn: "abs", x: 10.17, result: 10.17, ok: true},
		{fn: "abs", x: -10.17, result: 10.17, ok: true},
	}

	for _, tc := range tests {
		fn, ok := Function(tc.fn)
		result := fn(tc.x)
		if !Equals(result, tc.result) || ok != tc.ok {
			t.Errorf("fn(%v) = { %v, %v} != { %v, %v }", tc.x, result, ok, tc.result, tc.ok)
		}
	}
}
