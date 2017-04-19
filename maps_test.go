package main

import (
	"testing"

	"golang.org/x/tour/wc"
)

// type `go test -run LoopsAndFunctions` in the shell
func TestMaps(t *testing.T) {
	wc.Test(WordCount)
}
