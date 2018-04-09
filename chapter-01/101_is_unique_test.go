package chapter01

import (
	"testing"
)

func TestUniqueRunes(t *testing.T) {
	var tests = map[string]bool{
		"a":           true,
		"ab":          true,
		"abcdefghijk": true,
		"abcxyz":      true,
		"aa":          false,
		"abcc":        false,
		"abccde":      false,
	}

	t.Run("using_map", func(t *testing.T) {
		for key, expected := range tests {
			if actual := IsUnique(key); actual != expected {
				t.Errorf("Expected unique=%t for %s", expected, key)
			}
		}
	})

	t.Run("using_bitset", func(t *testing.T) {
		for key, expected := range tests {
			if actual := IsUniqueUsingBitset(key); actual != expected {
				t.Errorf("Expected unique=%t for %s", expected, key)
			}
		}
	})
}

//
// Implementation
//

type RuneSet map[rune]struct{}

func NewRuneSet() *RuneSet {
	return &RuneSet{}
}

// Add r to the set and return true; if already present in
// the set then return false
func (set *RuneSet) Add(r rune) bool {
	if _, exists := (*set)[r]; exists {
		return false
	}
	(*set)[r] = struct{}{}
	return true
}

func IsUnique(s string) bool {
	set := NewRuneSet()
	for _, r := range []rune(s) {
		if ok := set.Add(r); !ok {
			return false
		}
	}
	return true
}

// assumes only ascii
func IsUniqueUsingBitset(s string) bool {
	set := 0
	for _, ch := range []rune(s) {
		val := ch - 'a'
		mask := 1 << uint32(val)
		if set&mask > 0 {
			return false
		}
		set |= mask
	}
	return true
}
