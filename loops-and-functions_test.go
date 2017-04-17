package main

import (
	"fmt"
	"math"
	"testing"
)

// type `go test -run LoopsAndFunctions` in the shell
func TestLoopsAndFunctions(t *testing.T) {
	z := 1.0
	shouldBe2 := Sqrt(float64(4.0), z, 10)
	if shouldBe2 != 2.0 {
		t.Error("Sqrt(4) should be 2.0")
	}

	x := 5
	fmt.Println(Sqrt(float64(x), z, 1000))
	fmt.Println(math.Sqrt(float64(x)))
}
