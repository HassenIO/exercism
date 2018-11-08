// Package triangle is an exercise on exercism.io.
package triangle

import "math"

// Kind represents the kind of triangles.
type Kind string

const (
	// NaT is a kind of malformed triangle.
	NaT Kind = "NaT"
	// Equ is a kind of triangle where only two sides have equal lengths.
	Equ Kind = "Equ"
	// Iso is a kind of triangle where all sides have equal lengths.
	Iso Kind = "Iso"
	// Sca is a kind of triangle where all sides have different lengths.
	Sca Kind = "Sca"
)

// KindFromSides gives the type of triangle given length of all sides.
func KindFromSides(a, b, c float64) Kind {
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return NaT // Due to NaN value somewhere
	}
	if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return NaT // Due to Inf value somewhere
	}
	if a <= 0 || b <= 0 || c <= 0 || a+b < c || a+c < b || b+c < a {
		return NaT
	}
	if a == b && b == c {
		return Equ
	}
	if a == b && a != c || a == c && a != b || b == c && b != a {
		return Iso
	}
	return Sca
}
