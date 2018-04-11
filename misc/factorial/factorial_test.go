package factorial

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	tests := [][]int{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
		{6, 720},
	}

	for _, test := range tests {
		if actual := factorial(test[0]); actual != test[1] {
			t.Errorf("\nfor n=%d, expected: %d, actual: %d", test[0], test[1], actual)
		}
	}
}

// factorial is the product of an integer n and all integers less than it
// for n = 0 or n = 1, return 1
func factorial(n int) int {
	switch n {
	case 0: // base case
		fallthrough
	case 1: // base case
		return 1
	default:
		return n * factorial(n-1)
	}
}
