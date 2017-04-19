package main

import (
	"strings"
)

// WordCount count character of s and return count
func WordCount(s string) map[string]int {
	result := make(map[string]int)
	for _, value := range strings.Fields(s) {
		result[value] = result[value] + 1
	}
	return result
}
