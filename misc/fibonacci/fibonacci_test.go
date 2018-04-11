package fibonacci

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	type sample struct {
		n        int
		expected int
	}

	tests := []sample{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
	}

	t.Run("recursive", func(t *testing.T) {
		for _, samp := range tests {
			actual := fibRecursive(samp.n)
			expected := samp.expected

			if actual != expected {
				t.Errorf("\nfor %d, expected: %d, actual: %d", samp.n, expected, actual)
			}
		}
	})

	t.Run("best", func(t *testing.T) {
		for _, samp := range tests {
			actual := fib(samp.n)
			expected := samp.expected

			if actual != expected {
				t.Errorf("\nfor %d, expected: %d, actual: %d", samp.n, expected, actual)
			}
		}
	})
}

func fibRecursive(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fibRecursive(n-2) + fibRecursive(n-1)
	}
}

func fib(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	}

	a := 0
	b := 1
	sum := 0
	for i := 2; i <= n; i++ {
		sum = a + b
		a = b
		b = sum
	}

	return sum
}
