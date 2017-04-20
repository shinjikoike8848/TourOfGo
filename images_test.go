package main

import (
	"testing"

	"golang.org/x/tour/pic"
)

func TestImage(t *testing.T) {
	m := Image{100, 200}
	pic.ShowImage(m)
}
