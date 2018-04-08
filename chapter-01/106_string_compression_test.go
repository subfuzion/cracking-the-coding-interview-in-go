package chapter01

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringCompression(t *testing.T) {
	type sample struct {
		Input    string
		Expected string
	}

	tests := []sample{
		sample{
			"aabcccccaaa",
			"a2b1c5a3",
		},
		sample{
			"a",
			"a", // not a1
		},
		sample{
			"aa",
			"aa", // not a2
		},
		sample{
			"aaa",
			"a3",
		},
		sample{
			"hello world",
			"hello world", // not h1e1l2o1 1w1o1r1l1d1"
		},
	}

	for _, samp := range tests {
		actual := compress(samp.Input)
		if actual != samp.Expected {
			t.Errorf("sample: '%s', expected: '%s', actual: '%s'", samp.Input, samp.Expected, actual)
		}
	}
}

func compress(s string) string {
	// use a string builder for efficiency while building compressed string
	b := strings.Builder{}

	// string as a character (rune) array
	chars := []rune(s)

	// same character counter
	n := 0

	for i := 0; i < len(chars); i++ {
		n++
		if i+1 >= len(chars) || chars[i] != chars[i+1] {
			b.WriteRune(rune(chars[i]))
			b.WriteString(strconv.Itoa(n))
			n = 0
		}
	}

	// don't return compressed string unless it's actually shorter than the original
	if len(s) <= b.Len() {
		return s
	}

	return b.String()
}
