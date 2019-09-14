/*

FizzBuzz

fizzbuzz prints to a string all digits from 1 to n so that:
- for every 3rd digit, print "Fizz"
- for every 5 digit, print "Buzz
- for every multiple of 3 and 5, print "FizzBuzz"

Example:
  for n = 15
  result = "12Fizz4BuzzFizz78FizzBuzz11Fizz1314FizzBuzz"
			1
			 2
			  3 print "Fizz"
				  4
				   5 print "Buzz"
					   6 print "Fizz" (multiple of 3)
						   7
							8
							 9 print "Fizz" (multiple of 3)
								 10 print "Buzz" (multiple of 5)
									 11
									   12 print "Fizz" (multiple of 3)
										   13
											 14
											   15 print "FizzBuzz"
											      (multiple of 3 and 5)


*/

package misc

import (
	"strconv"
	"strings"
	"testing"
)

// FizzBuzzFunc declares a type of FizzBuzz function
type FizzBuzzFunc func(int) string

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

	test := func(t *testing.T, name string, f FizzBuzzFunc) {
		for _, samp := range tests {
			actual := f(samp.n)
			expected := samp.expected

			if actual != expected {
				t.Errorf("\nfor n=%d, expected: %s, actual: %s", samp.n, expected, actual)
			}
		}
	}

	test(t, "fizzbuzz", fizzbuzz)
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
