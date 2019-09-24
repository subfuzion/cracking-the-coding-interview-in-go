/*

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
package fibonacci

// Func declares a type of function such
// that Func(n) => the fibonacci number of n
type Func func(int) int

// TestCase for Func(input) defines expected result
type TestCase struct {
	Input    int
	Expected int
}

// SampleTestCases contains test data (maps PrimeFunc(input) => expected result)
var SampleTestCases = []TestCase{
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

//
// Solution
//
// Listed in order of increasing speed.
// There is a significant difference in performance between the recursive
// solutions (one memoizes).
//
// There is a significant difference in performance between the recursive
// and non-recursive solutions. The non-recursive solutions don't make
// function calls and won't run out of stack.
//
// Run benchmarks to analyze for yourself.
//

func FibRecursive1(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return FibRecursive1(n-2) + FibRecursive1(n-1)
	}
}

func FibRecursive2(n int) int {
	if n < 2 {
		return n
	}

	var memo = make([]int, n+1)
	memo[1] = 1

	var memoize func(int) int
	memoize = func(n int) int {
		if n < 2 {
			return n
		}

		// already memoized?
		// if memo[n] != 0 {
		// 	return memo[n]
		// }

		a := memo[n-2]
		if a == 0 {
			a = memoize(n - 2)
		}

		b := memo[n-1]
		if b == 0 {
			b = memoize(n - 1)
		}

		// make sure to memoize result
		memo[n] = a + b
		return memo[n]
	}
	return memoize(n-2) + memoize(n-1)
}

func Fib1(n int) int {
	if n < 2 {
		return n
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

func Fib2(n int) int {
	f := n

	a := 0
	b := 1
	for i := 2; i <= n; i++ {
		f = a + b
		a = b
		b = f
	}
	return f
}
