// Copyright 2018 Jeff Channell. All rights reserved.
// https://en.wikipedia.org/wiki/Golden_ratio

package golden

import "math"

// Ratio finds the golden ratio given two numbers a and b and returns c
// when a < b, a < c < b
// when a > b, c > a > b
// when a == b, a == b == c
func Ratio(a float64, b float64) (c float64) {
	// use interior division to find c
	// 1. Having a line segment AB, construct a perpendicular BC at point B, with BC half the length of AB. Draw the hypotenuse AC.
	ab := math.Abs(b - a)          // get the distance between the two "points" (a,0) and (b,0)
	bc := ab / 2                   // get the distance of the perpendicular (b,0) and (b,c)
	ac := math.Sqrt(ab*ab + bc*bc) // get the distance of the hypotenuse
	// 2. Draw an arc with center C and radius BC. This arc intersects the hypotenuse AC at point D.
	// 3. Draw an arc with center A and radius AD. This arc intersects the original line segment AB at point S. Point S divides the original line segment AB into line segments AS and SB with lengths in the golden ratio.
	c = a + (ac - bc) // instead of an arc, we know that BC is the radius so that can be subtracted from AC to get the distance of AD and to derive AS (it's the same distance as AD), adjusted for a
	return
}
