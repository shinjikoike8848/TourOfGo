package main

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	current, next := 1, 0
	return func() int {
		result := next
		next = current + next
		current = result
		return result
	}
}
