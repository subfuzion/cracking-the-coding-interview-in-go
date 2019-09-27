package prime

import (
	"testing"
)

//
// Tests
//
// $ go test -timeout 30s
//

// Tests - add additional test solutions here
var Tests = map[string]Func{
	"Prime1": Prime1,
	"Prime2": Prime2,
}

// TestSolution will run tests for the named solutions mapped
// in the tests var
func TestSolution(t *testing.T) {
	for test, f := range Tests {
		RunTest(t, test, f)
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
	for name, f := range Tests {
		n := 2147483647 // 64-bit: 18446744073709551557
		RunBenchmark(b, name, f, n)
	}
}
