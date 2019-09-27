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
	{10, 55},
	{11, 89},
	{12, 144},
	{13, 233},
	{14, 277},
	{15, 610},
}

//
// Solution
//
// Listed in order of increasing speed.
// There is a significant difference in performance between the non-recursive
// and recursive solutions,
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

func FibRecursiveMemo1(n int) int {
	if n < 2 {
		return n
	}

	memo := make([]int, n+1)
	memo[1] = 1

	var f func(int) int
	f = func(n int) int {
		if n < 2 {
			return n
		}
		if memo[n] == 0 {
			memo[n] = f(n-2) + f(n-1)
		}
		return memo[n]
	}

	return f(n-2) + f(n-1)
}

func FibRecursiveMemo2(n int) int {
	if n < 2 {
		return n
	}

	var fibs = make([]int, n)
	fibs[1] = 1

	var fib func(int) int

	fib = func(n int) int {
		if n == 0 {
			return 0
		}

		// already memoized?
		if fibs[n] != 0 {
			return fibs[n]
		}

		a := fibs[n-2]
		if a == 0 {
			a = fib(n - 2)
		}

		b := fibs[n-1]
		if b == 0 {
			b = fib(n - 1)
		}

		// make sure to memoize result
		fibs[n] = a + b
		return fibs[n]
	}

	return fib(n-2) + fib(n-1)
}

func Fib1(n int) int {
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

func Fib2(n int) int {
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

func Fib3(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
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

// Fibx1 uses only two array slots for space.
func Fibx1(n int) int {
	arr := []int{0, 1}

	for i := 2; i <= n; i++ {
		arr[i%2] = arr[0] + arr[1]
	}

	return arr[n%2]
}

// Fibx2 uses only two variables for space
func Fibx2(n int) int {
	if n < 2 {
		return n
	}

	// At the start of each loop:
	//   a = n-1
	//   b = n-2
	// At the end of each loop, a = fib(i)
	a := 1
	b := 0
	for i := 2; i <= n; i++ {
		a = a + b
		b = a - b
	}

	return a
}
