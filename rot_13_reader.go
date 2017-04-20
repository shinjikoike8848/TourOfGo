package main

import (
	"fmt"
	"io"
	//"os"
	//"strings"
)

type rot13Reader struct {
	r io.Reader
}

func isCapitalLetter(x byte) bool {
	return x >= 'A' && x <= 'Z'
}

func isSmallLetter(x byte) bool {
	return x >= 'a' && x <= 'z'
}

func (r13 *rot13Reader) Read(buffer []byte) (int, error) {
	fmt.Println(r13.r)
	length, err := r13.r.Read(buffer)
	for index := 0; index <= length; index++ {
		oneCharacter := buffer[index]
		if !(isCapitalLetter(oneCharacter) || isSmallLetter(oneCharacter)) {
			continue
		}

		encryptedValue := oneCharacter + 13
		if isCapitalLetter(oneCharacter) && encryptedValue > 'Z' || isSmallLetter(oneCharacter) && encryptedValue > 'z' {
			encryptedValue = encryptedValue - 26
		}
		buffer[index] = encryptedValue
	}
	return length, err
}
