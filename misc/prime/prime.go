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

// Func declares a type of prime number function such
// that Func(n) => true if n is prime
type Func func(int) bool

// TestCase for Func(input) defines expected result
type TestCase struct {
	Input    int
	Expected bool
}

// SampleTestCases contains test data (maps PrimeFunc(input) => expected result)
var SampleTestCases = []TestCase{
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

//
// Solution
//
// The solution isn't worth optimizing. The following are all 0(n).
// They are ranked in order of fastest to slowest, but the actual
// difference for Prime(10^10) was negligible between them.
//
// Run benchmarks on your own machine and see for yourself:
// go test -bench .

// This one is clear and technically was the fastest.
func Prime1(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// This one appears to be optimized and yet benchmarks reveal
// very little benefit.
func Prime2(n int) bool {
	mid := n / 2
	for i := 2; i <= mid; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n > 1
}

// Nice and concise...
// Compared to the clarity of the first solution, is the
// brevity of this one worth it? (For some types of solutions,
// the answer might be yes.)
func Prime3(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n > 1
}

// Nothing gained here by using a switch statement.
func Prime4(n int) bool {
	switch n {
	case 0:
		return false
	case 1: return false
	case 2: return true
	default:
		for i := 2; i < n; i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	}
}

