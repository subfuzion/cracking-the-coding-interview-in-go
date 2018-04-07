package chapter01

import (
	"strings"
	"testing"
)

func TestPalindromePermutation(t *testing.T) {
	type sample struct {
		Input      string
		Palindrome string
		Expected   bool
	}

	tests := []sample{
		sample{"atco cta", "taco cat", true},
		sample{"aan", "ana", true},
		sample{"anan", "anna", true},
		sample{"aaanx", "", false},
	}

	for _, samp := range tests {
		actual := isPalindromePermutation(samp.Input)
		expected := samp.Expected
		if actual != expected {
			if expected {
				t.Errorf("expected '%s' => '%s'", samp.Input, samp.Palindrome)
			} else {
				t.Errorf("unexpected: '%s'", samp.Input)
			}
		}
	}
}

func isPalindromePermutation(s string) bool {
	bv := createBitVector(s)
	return bv == 0 || onlyOneBitSet(bv)
}

func createBitVector(s string) uint32 {
	var bv uint32

	// assume ascii
	chars := strings.Split(s, "")
	for _, ch := range chars {
		x := getCharValue(ch)
		bv = toggle(bv, x)
	}
	return bv
}

func getCharValue(ch string) int {
	val := []rune(ch)[0]
	if 'a' <= val && val <= 'z' {
		return int(val - 'a')
	}
	return -1
}

func toggle(bv uint32, x int) uint32 {
	if x < 0 {
		return bv
	}

	var mask uint32 = 1 << uint32(x)
	if (bv & mask) == 0 {
		bv |= mask
	} else {
		bv &= ^mask
	}
	return bv
}

func onlyOneBitSet(bv uint32) bool {
	return (bv & (bv - uint32(1))) == 0
}
