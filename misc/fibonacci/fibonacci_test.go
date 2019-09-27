package fibonacci

import "testing"

//
// Tests
//
// $ go test -timeout 30s
//

// TODO: add additional test solutions here
var tests = map[string]Func{
	"FibRecursive1":     FibRecursive1,
	"FibRecursiveMemo1": FibRecursiveMemo1,
	"FibRecursiveMemo2": FibRecursiveMemo2,
	"Fib1":              Fib1,
	"Fib2":              Fib2,
	"Fib3":              Fib3,
	"Fibx1":             Fibx1,
	"Fibx2":             Fibx2,
}

// TestSolution runs tests for the named solutions mapped in the tests var
// using test data from SampleTestCases
func TestSolution(t *testing.T) {
	for test, f := range tests {
		t.Run(name, func(t *testing.T) {
			for _, test := range SampleTestCases {
				if actual := f(test.Input); actual != test.Expected {
					t.Errorf("\nfor n=%d, expected: %d, actual: %d", test.Input, test.Expected, actual)
				}
			}
		})
	}
}

//
// Benchmarks
//
// $ go test -benchmem -bench .
//

// BenchmarkSolution will run benchmarks for the named solutions mapped
// in the tests var
func BenchmarkSolution(b *testing.B) {
	for name, f := range tests {
		// TODO: pick a reasonable value for n to benchmark f(n)
		n := 20
		b.Run(name, func(b *testing.B) {
			// b.Logf("f(%d), loop (b.N) = %d\n", n, b.N)
			for i := 0; i < b.N; i++ {
				f(n)
			}
		})
	}
}
