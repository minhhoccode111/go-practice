// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides identify which type of triangle is formed.
func KindFromSides(a, b, c float64) Kind {
	switch {
	case !isTriangle(a, b, c):
		return NaT
	case isEquilateral(a, b, c):
		return Equ
	case isIsosceles(a, b, c):
		return Iso
	default:
		return Sca
	}
}

func isTriangle(a, b, c float64) bool {
	return a+b > c && b+c > a && a+c > b && a > 0 && b > 0 && c > 0
}

func isEquilateral(a, b, c float64) bool {
	return a == b && b == c
}

func isIsosceles(a, b, c float64) bool {
	return a == b || b == c || a == c
}
