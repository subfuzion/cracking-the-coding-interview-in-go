package fibonacci

import "testing"

//
// Tests
//
// $ go test -timeout 30s
//

// TODO: add additional test solutions here
var tests = map[string]Func{
	"FibRecursive1": FibRecursive1,
	"FibRecursive2": FibRecursive2,
	"Fib1": Fib1,
	"Fib2": Fib2,
	"Fib3": Fib3,
}

// TestSolution will run tests for the named solutions mapped
// in the tests var
func TestSolution(t *testing.T) {
	for test, f := range tests {
		RunTest(t, test, f, SampleTestCases)
	}
}

// RunTest runs a test using the supplied function and test data
func RunTest(t *testing.T, name string, f Func, testCases []TestCase) {
	t.Run(name, func(t *testing.T) {
		for _, test := range testCases {
			if actual := f(test.Input); actual != test.Expected {
				t.Errorf("\nfor n=%d, expected: %d, actual: %d", test.Input, test.Expected, actual)
			}
		}
	})
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
		RunBenchmark(name, b, f, n)
	}
}

// RunBenchmark runs a benchmark using the supplied function and test data
func RunBenchmark(name string, b *testing.B, f Func, n int) {
	b.Run(name, func(b *testing.B) {
		b.Logf("f(%d), loop (b.N) = %d\n", n, b.N)
		for i := 0; i < b.N; i++ {
			f(n)
		}
	})
}
