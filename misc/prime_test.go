/*

Prime

Prime is a whole number n that can only be divided by 1 and itself:
- for n = 0, return false
- for n = 1, return false
- for all other n, return true if its only two factors are 1 and n

Examples:
0 =>  false
1 =>  false
2 =>  true
3 =>  true
4 =>  false (other factors: 2)
5 =>  true
6 =>  false (other factors: 2, 3)
7 =>  true
8 =>  false (other factors: 2, 4)
9 =>  false (other factors: 3)
10 => false (other factors: 2, 5)
11 => true
30 => false (other factors: 2, 3, 5, 6, 10, 15)

Note: 2 is the only even number that is prime. No other even number n
can be prime since, in addition to factors (1, n), there will at least
be factors (2, n/2).

*/

package misc

import (
	"testing"
)

// PrimeFunc declares a type of prime number function
type PrimeFunc func(int) bool

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

	test := func(t *testing.T, name string, f PrimeFunc) {
		for _, samp := range tests {
			actual := prime(samp.n)
			expected := samp.expected

			if expected != actual {
				t.Errorf("\nfor n=%d, expected: %t, actual: %t", samp.n, expected, actual)
			}
		}
	}

	test(t, "prime", prime)
}

func prime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n > 1
}
