package main

import (
	"fmt"
	"math"
	"testing"
)

const sqrt4 = 2.0

// type `go test -run LoopsAndFunctions` in the shell
func TestSqrt(t *testing.T) {
	shouldBe2 := Sqrt(float64(4.0))
	fmt.Println(shouldBe2)

	if shouldBe2 != sqrt4 {
		t.Errorf("Sqrt(4) should be %f but accepted value is %f", sqrt4, shouldBe2)
	}

	x := 5
	fmt.Println(Sqrt(float64(x)))
	fmt.Println(math.Sqrt(float64(x)))
}
