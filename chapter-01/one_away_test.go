package chapter01

import (
	"testing"
)

func TestOneAway(t *testing.T) {
	type sample struct {
		A        string
		B        string
		Category string
		Expected bool
	}

	tests := []sample{
		sample{
			"abc",
			"abd",
			"one replacement",
			true,
		},
		{
			"abc",
			"abb",
			"one replacement",
			true,
		},
		{
			"bbc",
			"abc",
			"one replacement",
			true,
		},
		{
			"abc",
			"abcd",
			"one insert",
			true,
		},
		{
			"ac",
			"abc",
			"one insert",
			true,
		},
		{
			"bc",
			"abc",
			"one insert",
			true,
		},
		{
			"abc",
			"bc",
			"one insert",
			true,
		},
		{
			"abc",
			"ac",
			"one insert",
			true,
		},
		{
			"abc",
			"ab",
			"one insert",
			true,
		},
		{
			"ab",
			"abcd",
			"more than one away insert",
			false,
		},
		{
			"abcde",
			"ab",
			"more than one away insert",
			false,
		},
		{
			"abcd",
			"abbb",
			"more than one away replace",
			false,
		},
		{
			"abc",
			"xyz",
			"more than one away replace",
			false,
		},
		{
			"abc",
			"abc",
			"equal strings",
			true,
		},
	}

	for _, samp := range tests {
		if ok := isOneEditAway(samp.A, samp.B); ok != samp.Expected {
			t.Errorf("%s: '%s', '%s'", samp.Category, samp.A, samp.B)
		}
	}
}

func isOneEditAway(a, b string) bool {
	if a == b {
		return true
	}

	d := len(a) - len(b)
	switch d {
	case 0:
		return oneEditReplace(a, b)
	case -1:
		return oneEditInsert(a, b)
	case 1:
		return oneEditInsert(b, a)
	default:
		// difference is too great for a single edit
		return false
	}
}

func oneEditReplace(a, b string) bool {
	found := false
	ra := []rune(a)
	rb := []rune(b)

	for i := range ra {
		if ra[i] != rb[i] {
			// if found a mismatch already, then this is more than one edit away
			if found {
				return false
			}
			found = true
		}
	}
	return found
}

func oneEditInsert(a, b string) bool {
	ra := []rune(a)
	rb := []rune(b)

	for ia, ib := 0, 0; ib < len(ra); {
		if ra[ia] == rb[ib] {
			ia++
			ib++
		} else {
			ib++
		}

		if ib-ia > 1 {
			return false
		}
	}

	return true
}
