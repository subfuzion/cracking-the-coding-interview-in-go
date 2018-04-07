package chapter01

import (
	"testing"
)

func TestStringRotation(t *testing.T) {
	type sample struct {
		A        string
		B        string
		Expected bool
	}

	tests := []sample{
		sample{
			"a",
			"a",
			true,
		},
		sample{
			"ab",
			"ab",
			true,
		},
		sample{
			"ab",
			"ba",
			true,
		},
		sample{
			"abc",
			"bca",
			true,
		},
		sample{
			"abc",
			"cab",
			true,
		},
		sample{
			"waterbottle",
			"erbottlewat",
			true,
		},
		sample{
			"abc",
			"cba",
			false,
		},
		sample{
			"",
			"",
			false,
		},
		sample{
			"ab",
			"abc",
			false,
		},
		sample{
			"abc",
			"xyz",
			false,
		},
	}

	for _, samp := range tests {
		actual := isStringRotation(samp.A, samp.B)
		expected := samp.Expected

		if actual != expected {
			t.Errorf("'%s' is a rotation of '%s' should be %t", samp.B, samp.A, samp.Expected)
		}
	}
}

func isStringRotation(a, b string) bool {
	n := len(a)
	if n == 0 || n != len(b) {
		return false
	}

	aa := a + a
	// check that b is a substring of aa (as will always be the case if b is a valid rotation)
	return isSubstring(aa, b)
}

func isSubstring(a, b string) bool {
	return index(a, b) >= 0
}

// see https://golang.org/src/strings/strings_amd64.go
// for industrial strength solution
func index(a, b string) int {
	n := len(b)
	switch {
	case n == 0:
		return 0
	case n == len(a):
		if b == a {
			return 0
		}
		return -1
	case n > len(a):
		return -1
	}

	ra := []rune(a)
	rb := []rune(b)

	for i := 0; i < len(ra)-len(rb); i++ {
		for j := 0; j < len(rb); j++ {
			if rb[j] != ra[i+j] {
				break
			}
			if j+1 == len(rb) {
				return i
			}
		}
	}

	return -1
}
