/*

Prime is a whole number n that can only be divided by 1 and itself:
- for n = 0, return false
- for n = 1, return false
- for all other n, return true if its only two factors are 1 and n

Examples:
0   =>  false
1   =>  false
2   =>  true
3   =>  true
4   =>  false (other factors: 2)
5   =>  true
6   =>  false (other factors: 2, 3)
7   =>  true
8   =>  false (other factors: 2, 4)
9   =>  false (other factors: 3)
10  =>  false (other factors: 2, 5)
11  =>  true
30  =>  false (other factors: 2, 3, 5, 6, 10, 15)

Note: 2 is the only even number that is prime. No other even number n
can be prime since, in addition to factors (1, n), there will at least
be factors (2, n/2).

*/
package prime

import "testing"

// Func declares a type of prime number function such
// that Func(n) => true if n is prime
type Func func(int) bool

// Test for Func(input) defines expected result
type Test struct {
	Input    int
	Expected bool
}

// TestData contains test data (maps PrimeFunc(input) => expected result)
var TestData = []Test{
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
	{12, false},
	{13, true},
	{14, false},
	{15, false},
	{16, false},
	{17, true},
	{18, false},
	{19, true},
	{20, false},
	{21, false},
	{22, false},
	{23, true},
	{24, false},
	{25, false},
	{26, false},
	{27, false},
	{28, false},
	{29, true},
	{30, false},
	{70, false},
	{71, true},
	{72, false},
	{228, false},
	{229, true},
	{230, false},
}

// RunTest runs a test using the supplied function and test data
func RunTest(t *testing.T, name string, f Func) {
	t.Run(name, func(t *testing.T) {
		for _, test := range TestData {
			if actual := f(test.Input); actual != test.Expected {
				t.Errorf("\nfor n=%d, expected: %t, actual: %t", test.Input, test.Expected, actual)
			}
		}
	})
}

// RunBenchmark runs a benchmark using the supplied function and argument
func RunBenchmark(b *testing.B, name string, f Func, n int) {
	b.Run(name, func(b *testing.B) {
		// b.Logf("f(%d), loop (b.N) = %d\n", n, b.N)
		for i := 0; i < b.N; i++ {
			f(n)
		}
	})
}
