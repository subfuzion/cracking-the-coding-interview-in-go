package chapter01

import (
	"sort"
	"strings"
	"testing"
)

func TestPermutations(t *testing.T) {
	type sample struct {
		A        string
		B        string
		Expected bool
	}

	tests := []sample{
		sample{A: "abc", B: "abc", Expected: true},
		sample{A: "abc", B: "bca", Expected: true},
		sample{A: "abc", B: "cab", Expected: true},
		sample{A: "abc", B: "ab", Expected: false},
		sample{A: "abc", B: "abcd", Expected: false},
		sample{"a", "a", true},
		sample{"a", "b", false},
		sample{"", "b", false},
		sample{"a", "", false},
	}

	t.Run("by sorting", func(t *testing.T) {
		for _, samp := range tests {
			if actual := isPermuationBySort(samp.A, samp.B); actual != samp.Expected {
				t.Errorf("expected %t for a=%s, b=%s", samp.Expected, samp.A, samp.B)
			}
		}
	})

	t.Run("by weight", func(t *testing.T) {
		for _, samp := range tests {
			if actual := isPermuationByWeight(samp.A, samp.B); actual != samp.Expected {
				t.Errorf("expected %t for a=%s, b=%s", samp.Expected, samp.A, samp.B)
			}
		}
	})
}

func isPermuationBySort(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	// assume ascii
	achars := strings.Split(a, "")
	bchars := strings.Split(b, "")

	sort.Strings(achars)
	sort.Strings(bchars)

	sa := strings.Join(achars, "")
	sb := strings.Join(bchars, "")

	return sa == sb
}

func isPermuationByWeight(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	weights := map[rune]int{}

	// for each rune in a, increment its count
	for _, r := range a {
		weights[r]++
	}

	// for each rune in b, decrement its count
	// if we go negative, then there was no corresponding rune
	// in a and we can short circuit out with false
	for _, r := range b {
		weights[r]--
		if weights[r] < 0 {
			return false
		}
	}

	// since the strings were the same length and all letters were accounted
	// for, we can return with success
	return true
}
