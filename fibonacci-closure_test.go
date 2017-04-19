package main

import (
	"fmt"
	"testing"
)

func TestFibonacciClosure(t *testing.T) {
	fmt.Println("fibonacci closure")
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
