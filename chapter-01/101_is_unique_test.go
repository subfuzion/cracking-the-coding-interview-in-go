package chapter01

import (
	"testing"
)

func TestUniqueRunes(t *testing.T) {
	var tests = map[string]bool{
		"a":           true,
		"ab":          true,
		"abcdefghijk": true,
		"abc123@=%^":  true,
		"aa":          false,
		"abcc":        false,
		"abccde":      false,
	}

	for key, expected := range tests {
		if actual := IsUnique(key); actual != expected {
			t.Errorf("Expected unique=%t for %s", expected, key)
		}
	}
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
	runes := []rune{}
	for _, r := range s {
		runes = append(runes, r)
	}

	set := NewRuneSet()
	for _, r := range runes {
		if ok := set.Add(r); !ok {
			return false
		}
	}
	return true
}
