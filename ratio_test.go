// Copyright 2018 Jeff Channell. All rights reserved.
package golden

import (
	"math"
	"testing"
)

func TestRatio(t *testing.T) {
	tests := []struct {
		a float64
		b float64
		e float64
	}{
		// tests where a < b
		{a: 0, b: 8, e: 5},
		{a: 0, b: 13, e: 8},
		{a: 3, b: 16, e: 11},
		{a: -3, b: 10, e: 5},
		// tests where a > b
		{a: 5, b: 0, e: 8},
		{a: 8, b: 0, e: 13},
		// tests where a == b
		{a: 0, b: 0, e: 0},
		{a: 1, b: 1, e: 1},
		{a: -1, b: -1, e: -1},
	}
	for i, test := range tests {
		f := Ratio(test.a, test.b)
		if r := int(math.Round(f)); int(test.e) != r {
			t.Fatalf("case %d: Expected %v but got %d; %v", i, test.e, r, f)
		}
	}
}
