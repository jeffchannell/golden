// Copyright 2018 Jeff Channell. All rights reserved.
// https://en.wikipedia.org/wiki/Golden_ratio

package golden

import (
	"math"
)

// Next finds the number that satifies the golden ratio using the length of a segment AS for the base
func Next(a float64, s float64) (b float64) {
	// Dividing a line segment by exterior division
	// 1. Draw a line segment AS and construct off the point S a segment SC perpendicular to AS and with the same length as AS.
	// get the distance between the two "points" (a,0) and (b,0)
	as := math.Abs(a - s)
	// 2. Do bisect the line segment AS with M.
	am := as / 2
	// 3. A circular arc around M with radius MC intersects in point B the straight line through points A and S (also known as the extension of AS). The ratio of AS to the constructed segment SB is the golden ratio.
	// instead of an arc, we know that the value is the length of CM + AM, so find CM
	// substitute AS for CS and AM for SM in these calculations, as they are the same sizes
	// then offset with either a or s based on size
	cm := math.Sqrt(as*as + am*am)
	if a > s {
		b = s - (cm - am)
	} else {
		b = a + (cm + am)
	}
	return
}

// Ratio finds the golden ratio between two numbers a and b
func Ratio(a float64, b float64) (c float64) {
	// use interior division to find c
	// 1. Having a line segment AB, construct a perpendicular BC at point B, with BC half the length of AB. Draw the hypotenuse AC.
	// get the distance between the two "points" (a,0) and (b,0)
	pos := (a > b)
	ab := math.Abs(a - b)
	bc := ab / 2                   // get the distance of the perpendicular (b,0) and (b,c)
	ac := math.Sqrt(ab*ab + bc*bc) // get the distance of the hypotenuse
	// 2. Draw an arc with center C and radius BC. This arc intersects the hypotenuse AC at point D.
	// 3. Draw an arc with center A and radius AD. This arc intersects the original line segment AB at point S. Point S divides the original line segment AB into line segments AS and SB with lengths in the golden ratio.
	// instead of an arc, we know that BC is the radius so that can be subtracted from AC to get the distance of AD and to derive AS (it's the same distance as AD), adjusted for a
	if pos {
		c = b - (bc - ac)
	} else {
		c = b + (bc - ac)
	}
	return
}
