package prime

import (
	"testing"
)

func TestPrime(t *testing.T) {
	type sample struct {
		n        int
		expected bool
	}

	tests := []sample{
		{0, false},
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{7, true},
		{8, false},
		{9, false},
		{10, false},
		{11, true},
		{30, false},
	}

	for _, samp := range tests {
		actual := prime(samp.n)
		expected := samp.expected

		if expected != actual {
			t.Errorf("\nfor n=%d, expected: %t, actual: %t", samp.n, expected, actual)
		}
	}
}

func prime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n > 1
}
