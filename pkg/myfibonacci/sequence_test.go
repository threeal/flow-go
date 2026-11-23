package myfibonacci_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/threeal/go-starter/pkg/myfibonacci"
)

func TestFibonacciSequence(t *testing.T) {
	t.Parallel()

	assert.Equal(t, []int{}, myfibonacci.FibonacciSequence(-1))
	assert.Equal(t, []int{}, myfibonacci.FibonacciSequence(0))
	assert.Equal(t, []int{1}, myfibonacci.FibonacciSequence(1))
	assert.Equal(t, []int{1, 1}, myfibonacci.FibonacciSequence(2))
	assert.Equal(t, []int{1, 1, 2, 3, 5}, myfibonacci.FibonacciSequence(5))
}
