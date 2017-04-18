package main

// Sqrt return Sqrt value
// x: sqrted value
// z: z
// count: loop count
func Sqrt(x float64, z float64, count int) float64 {
	if count <= 1 {
		return z
	}
	z = z - (z*z-x)/(2.0*z)
	count--
	return Sqrt(x, z, count)
}
