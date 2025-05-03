package complexnumbers

import "math"

// Number represents a complex number with real and imaginary parts
type Number struct {
	a float64 // real part
	b float64 // imaginary part
}

// Real returns the real part of the complex number
func (n Number) Real() float64 {
	return n.a
}

// Imaginary returns the imaginary part of the complex number
func (n Number) Imaginary() float64 {
	return n.b
}

// Add adds two complex numbers
func (n1 Number) Add(n2 Number) Number {
	return Number{
		a: n1.a + n2.a,
		b: n1.b + n2.b,
	}
}

// Subtract subtracts two complex numbers
func (n1 Number) Subtract(n2 Number) Number {
	return Number{
		a: n1.a - n2.a,
		b: n1.b - n2.b,
	}
}

// Multiply multiplies two complex numbers
func (n1 Number) Multiply(n2 Number) Number {
	return Number{
		a: n1.a*n2.a - n1.b*n2.b,
		b: n1.a*n2.b + n1.b*n2.a,
	}
}

// Times multiplies a complex number by a real factor
func (n Number) Times(factor float64) Number {
	return Number{
		a: n.a * factor,
		b: n.b * factor,
	}
}

// Divide divides two complex numbers
func (n1 Number) Divide(n2 Number) Number {
	denominator := n2.a*n2.a + n2.b*n2.b
	return Number{
		a: (n1.a*n2.a + n1.b*n2.b) / denominator,
		b: (n1.b*n2.a - n1.a*n2.b) / denominator,
	}
}

// Conjugate returns the complex conjugate of a complex number
func (n Number) Conjugate() Number {
	return Number{
		a: n.a,
		b: -n.b,
	}
}

// Abs returns the absolute value (magnitude) of a complex number
func (n Number) Abs() float64 {
	return math.Sqrt(n.a*n.a + n.b*n.b)
}

// Exp returns e raised to the power of the complex number
func (n Number) Exp() Number {
	// e^(a + bi) = e^a * (cos(b) + i*sin(b))
	ea := math.Exp(n.a)
	return Number{
		a: ea * math.Cos(n.b),
		b: ea * math.Sin(n.b),
	}
}
