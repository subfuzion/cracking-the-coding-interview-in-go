package fizzbuzz

import (
	"strconv"
	"strings"
	"testing"
)

func TestFizzBuzz(t *testing.T) {

	type sample struct {
		n        int
		expected string
	}

	tests := []sample{
		sample{
			n:        15,
			expected: "12Fizz4BuzzFizz78FizzBuzz11Fizz1314FizzBuzz",
		},
	}

	for _, samp := range tests {
		actual := fizzbuzz(samp.n)
		expected := samp.expected

		if actual != expected {
			t.Errorf("\nfor n=%d, expected: %s, actual: %s", samp.n, expected, actual)
		}
	}
}

// fizzbuzz prints to a string all digits from 1 to n
// for every 3rd digit, print "Fizz"
// for every 5 digit, print "Buzz
// for every multiple of 3 and 5, print "FizzBuzz"
func fizzbuzz(n int) string {
	b := strings.Builder{}
	for i := 1; i <= n; i++ {
		switch {
		case i%(3*5) == 0:
			b.WriteString("FizzBuzz")
		case i%3 == 0:
			b.WriteString("Fizz")
		case i%5 == 0:
			b.WriteString("Buzz")
		default:
			b.WriteString(strconv.Itoa(i))
		}

	}
	return b.String()
}
