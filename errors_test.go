package main

import (
	"testing"
)

func TestSqrtWithError(t *testing.T) {
	positiveValue := 2.0
	negativeValue := -2.0

	_, err := SqrtWithError(positiveValue)
	if err != nil {
		t.Errorf("SqrtWithError return error object Even though It accept positive number(%f) as a argument", positiveValue)
	}

	_, err = SqrtWithError(negativeValue)
	if err == nil {
		t.Errorf("SqrtWithError return error object as nil Even though It accept negative number(%f) as a argument", negativeValue)
	}
}
