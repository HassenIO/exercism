// Package triangle is an exercise on exercism.io.
package triangle

import "math"

// Kind represents the kind of triangles.
type Kind int

const (
	// NaT is a kind of malformed triangle.
	NaT Kind = iota
	// Equ is a kind of triangle where only two sides have equal lengths.
	Equ
	// Iso is a kind of triangle where all sides have equal lengths.
	Iso
	// Sca is a kind of triangle where all sides have different lengths.
	Sca
)

// KindFromSides gives the type of triangle given length of all sides.
func KindFromSides(a, b, c float64) Kind {
	if math.IsNaN(a+b+c) || math.IsInf(a+b+c, 0) {
		return NaT // Due to some NaN or Inf values
	}
	if a <= 0 || b <= 0 || c <= 0 || a+b < c || a+c < b || b+c < a {
		return NaT
	}
	if a == b && b == c {
		return Equ
	}
	if a == b || a == c || b == c {
		return Iso
	}
	return Sca
}
