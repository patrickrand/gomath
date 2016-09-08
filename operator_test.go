package main

import (
	"math"
	"testing"
)

func TestOperator(t *testing.T) {
	tests := []struct {
		op     string
		a, b   float64
		result float64
		ok     bool
	}{
		{op: "+", a: 3, b: 10.17, result: 13.17, ok: true},
		{op: "-", a: 3, b: 10.17, result: -7.17, ok: true},
		{op: "*", a: 3, b: 10.17, result: 30.51, ok: true},
		{op: "/", a: 3, b: 10.17, result: 0.29498525, ok: true},
		{op: "/", a: 3, b: 0, result: math.Inf(1), ok: true},
		{op: "/", a: -3, b: 0, result: math.Inf(-1), ok: true},
		{op: "/", a: 0, b: 0, result: math.NaN(), ok: true},
		{op: "pow", a: 3, b: 10.17, result: 71174.29278084412, ok: true},
		{op: "?", a: 3, b: 10.17, result: 0, ok: false},
		{op: "!", a: 3, b: 10.17, result: 0, ok: false},
		{op: "x", a: 3, b: 10.17, result: 0, ok: false},
		{op: "%", a: 3, b: 10.17, result: 0, ok: false},
		{op: "|", a: 3, b: 10.17, result: 0, ok: false},
		{op: "(", a: 3, b: 10.17, result: 0, ok: false},
		{op: ")", a: 3, b: 10.17, result: 0, ok: false},
		{op: "=", a: 3, b: 10.17, result: 0, ok: false},
		{op: "~", a: 3, b: 10.17, result: 0, ok: false},
		{op: "'", a: 3, b: 10.17, result: 0, ok: false},
	}

	for _, tc := range tests {
		op, ok := Operator(tc.op)
		result := op(tc.a, tc.b)
		if !Equals(result, tc.result) || ok != tc.ok {
			t.Errorf("op(%v %s %v) = { %v, %v } != { %v, %v }", tc.a, string(tc.op), tc.b, result, ok, tc.result, tc.ok)
		}
	}
}
