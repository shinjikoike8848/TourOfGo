package main

import "fmt"

// ErrNegativeSqrt implements error interface
type ErrNegativeSqrt float64

func (negativeSqrt ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number: ", float64(negativeSqrt))
}

//SqrtWithError is the func for exercise error
func SqrtWithError(x float64) (float64, error) {
	z := sqrtBase
	var err error
	if x < 0 {
		return z, ErrNegativeSqrt(x)
	}
	for i := 10; i > 0; i-- {
		z = z - (z*z-x)/(2.0*z)
	}
	return z, err
}
