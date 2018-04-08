package chapter01

import (
	"testing"
	"unicode/utf8"
)

func TestURLify(t *testing.T) {
	type sample struct {
		Input    string
		Expected string
	}

	tests := []sample{
		sample{"hello world", "hello%20world"},
		sample{"how now, brown cow", "how%20now,%20brown%20cow"},
	}

	for _, samp := range tests {
		// for this exercise we are supposed to work with character arrays
		// we can assume that we have a character array large enough to hold the
		// correct result; so we will use the actual answer to determine the size
		// needed, and then populate it with the input
		nRunes := utf8.RuneCountInString(samp.Expected)
		runes := make([]rune, nRunes)
		for i, r := range samp.Input {
			runes[i] = r
		}

		urlify(runes, utf8.RuneCountInString(samp.Input))
		actual := string(runes)
		expected := samp.Expected

		if actual != expected {
			t.Errorf("actual: '%s', expected: '%s'", actual, expected)
		}

	}
}

// assumes rune array is large enough to hold the extra characters after conversion
func urlify(runes []rune, nRunes int) {
	ri := len(runes) - 1
	for i := nRunes - 1; i >= 0; i-- {
		if runes[i] == ' ' {
			runes[ri-2] = '%'
			runes[ri-1] = '2'
			runes[ri] = '0'
			ri -= 3
		} else {
			runes[ri] = runes[i]
			ri--
		}
	}
}
