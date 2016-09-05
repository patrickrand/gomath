package float

import "testing"

func TestEquals(t *testing.T) {
	tests := []struct {
		a, b   float64
		result bool
	}{
		{a: 0, b: Epsilon, result: false},
		{a: 0, b: Epsilon / 10, result: true},
		{a: -123, b: -(123 + Epsilon*10), result: false},
		{a: -123, b: -(123 + Epsilon), result: true},
	}

	for _, tc := range tests {
		if result := Equals(tc.a, tc.b); result != tc.result {
			t.Errorf("Equals(%v, %v) = %v != %v", tc.a, tc.b, result, tc.result)
		}
	}
}
