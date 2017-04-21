package main

import (
	"fmt"
	"testing"

	"golang.org/x/tour/tree"
)

func TestBinaryTree(t *testing.T) {
	treeInstance := tree.New(10)
	ch := make(chan int)
	go func() {
		Walk(treeInstance, ch)
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}

	k := 10
	if !Same(tree.New(k), tree.New(k)) {
		t.Errorf("Same func should return true when compare tree.New(sameValue)")
	}

	defK := 11
	if Same(tree.New(k), tree.New(defK)) {
		t.Errorf("Same func should return false when compare tree.New(defferenceValue)")
	}
}
