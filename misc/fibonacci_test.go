/*

Fibonacci

The fibonacci number of a non-negative integer n is defined as:
- for n = 0, return 0
- for n = 1, return 1
- for all other positive integers, return the sum of the fibonacci
  values for the two preceding integers

Examples:
0 =>  0 (by definition)
1 =>  1 (by definition)
2 =>  1
3 =>  2 = fib(2) + fib(1) =  1 +  1
4 =>  3 = fib(3) + fib(2) =  2 +  1
5 =>  5 = fib(4) + fib(3) =  3 +  2
6 =>  8 = fib(5) + fib(4) =  5 +  3
7 => 13 = fib(6) + fib(5) =  8 +  5
8 => 21 = fib(7) + fib(6) = 13 +  8
9 => 34 = fib(8) + fib(7) = 21 + 13

*/

package misc

import (
	"testing"
)

// FibonacciFunc declares a type of Fibonacci function
type FibonacciFunc func(int) int

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
		{9, 34},
	}

	test := func(t *testing.T, name string, fib FibonacciFunc) {
		t.Run(name, func(t *testing.T) {
			for _, samp := range tests {
				actual := fib(samp.n)
				expected := samp.expected

				if actual != expected {
					t.Fatalf("\nfor %d, expected: %d, actual: %d", samp.n, expected, actual)
				}
			}
		})
	}

	test(t, "recursive", fibRecursive)
	test(t, "non-recursive", fib)
}

// fibonacci value of a number n is equal to the sum of the fibonacci value of the two preceding numbers
// for n = 0, return 0; for n = 1, return 1
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

// fibonacci value of a number n is equal to the sum of the fibonacci value of the two preceding numbers
// for n = 0, return 0; for n = 1, return 1
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
