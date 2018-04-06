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
	}

	for _, samp := range tests {
		actual := compress(samp.Input)
		if actual != samp.Expected {
			t.Errorf("sample: '%s', expected: '%s', actual: '%s'", samp.Input, samp.Expected, actual)
		}
	}
}

func compress(s string) string {
	// string as a character (rune) array
	chars := []rune(s)

	// previous character to compare
	p := chars[0]

	// same character counter
	n := 1

	b := strings.Builder{}

	for i := 1; i < len(chars); i++ {
		r := chars[i]

		if r == p {
			n++
		} else {
			b.WriteRune(rune(p))
			b.WriteString(strconv.Itoa(n))
			n = 1
		}

		p = r
	}

	// don't forget to write out final character+count
	b.WriteRune(rune(p))
	b.WriteString(strconv.Itoa(n))

	if len(s) <= b.Len() {
		return s
	}

	return b.String()
}
