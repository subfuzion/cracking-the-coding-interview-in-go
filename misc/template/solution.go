/*

Foo is a function that for given input n ...
- base case: for n = ?, return ?
- for all other n, return ?

Examples:
0   =>  ?
1   =>  ?
50  =>
100 =>  ?

*/
package template

// Func declares a type of function such
// that Func(n) => ?
type Func func(int) bool

// TestCase for Func(input) defines expected result
type TestCase struct {
	Input    int
	Expected bool
}

// SampleTestCases contains test data (maps PrimeFunc(input) => expected result)
var SampleTestCases = []TestCase{
	{0, false},
}


//
// Solution
//

func Solution(n int) bool {
	return false
}
