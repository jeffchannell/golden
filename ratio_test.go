// Copyright 2018 Jeff Channell. All rights reserved.
package golden

import (
	"math"
	"testing"
)

func TestNext(t *testing.T) {
	tests := []struct {
		a float64
		c float64
		e float64
	}{
		// tests where a < c
		{a: 0, c: 3, e: 5},
		{a: 0, c: 5, e: 8},
		{a: 0, c: 8, e: 13},
		{a: 3, c: 8, e: 11},
		{a: -8, c: -3, e: 0},
		{a: -3, c: 2, e: 5},
		// tests where a > c
		{a: 3, c: 0, e: -2},
		{a: 5, c: 0, e: -3},
		{a: 8, c: 0, e: -5},
		{a: 16, c: 8, e: 3},
		{a: 5, c: -3, e: -8},
		{a: -3, c: -8, e: -11},
		// tests where a == c
		{a: 0, c: 0, e: 0},
		{a: 1, c: 1, e: 1},
		{a: -1, c: -1, e: -1},
	}
	// loop through the tests
	for i, test := range tests {
		// get the ratio
		c := Next(test.a, test.c)
		// check the expected result against the actual one
		if r := int(math.Round(c)); int(test.e) != r {
			t.Fatalf("case %d: Expected %v but got %d; %v", i, test.e, r, c)
		}
	}
}

func TestRatio(t *testing.T) {
	tests := []struct {
		a float64
		b float64
		e float64
	}{
		// tests where a < b
		{a: 0, b: 5, e: 2},
		{a: 0, b: 8, e: 3},
		{a: 0, b: 13, e: 5},
		{a: 3, b: 16, e: 8},
		{a: -3, b: 10, e: 2},
		// tests where a > b
		{a: 5, b: 0, e: 3},
		{a: 8, b: 0, e: 5},
		{a: 13, b: 0, e: 8},
		{a: 16, b: 3, e: 11},
		{a: 10, b: -3, e: 5},
		// tests where a == b
		{a: 0, b: 0, e: 0},
		{a: 1, b: 1, e: 1},
		{a: -1, b: -1, e: -1},
	}
	// loop through the tests
	for i, test := range tests {
		// get the ratio
		c := Ratio(test.a, test.b)
		// check the expected result against the actual one
		if r := int(math.Round(c)); int(test.e) != r {
			t.Fatalf("case %d: Expected %v but got %d; %v", i, test.e, r, c)
		}
	}
}
