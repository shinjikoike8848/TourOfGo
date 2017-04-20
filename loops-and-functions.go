package main

const base = 1.0

// Sqrt return calc Sqrt value of x
func Sqrt(x float64) float64 {
	z := base
	for i := 10; i > 0; i-- {
		z = z - (z*z-x)/(2.0*z)
	}
	return z
}
