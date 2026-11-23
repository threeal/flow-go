// Package myfibonacci provides example functions for generating sequences of numbers.
package myfibonacci

// FibonacciSequence generates a Fibonacci sequence up to the given number of terms.
func FibonacciSequence(terms int) []int {
	if terms <= 0 {
		return []int{}
	} else if terms == 1 {
		return []int{1}
	}

	sequence := make([]int, terms)
	sequence[0] = 1
	sequence[1] = 1

	for i := 2; i < terms; i++ {
		sequence[i] = sequence[i-2] + sequence[i-1]
	}

	return sequence
}
