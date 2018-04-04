package isunique

import (
	"strings"
	"testing"
)

//
// Test
//

var tests = map[string]bool{
	"a":        true,
	"abcdef":   true,
	"aa":       false,
	"abcc":     false,
	"abcdefgh": true,
}

func Test(t *testing.T) {
	s := "abcdef"
	if !isUnique(s) {
		t.Errorf("Not unique: '%s'", s)
	}

	for str, unique := range tests {
		if isUnique(str) != unique {
			t.Errorf("Expected '%s' = %t", str, unique)
		}
	}
}

//
// Implementation
//

func isUnique(s string) bool {
	// use empty space to split by UTF-8 characters
	// for a more robust solution, consider splitting on runes
	chars := strings.Split(s, "")
	set := NewStringSet()
	for _, char := range chars {
		// ok will be false if the character is already a set member
		if ok := set.Add(char); !ok {
			return false
		}
	}

	return true
}

// Set is the interface for an unordered set of data
type Set interface {
	Add(i interface{}) bool
	Remove(i interface{}) bool
	Contains(i ...interface{}) bool
}

// stringSet is a private type used to implement the Set interface for strings
// map of string to the empty struct{}
// (the emptry struct holds no data so is the optimal value
// to use to indicate string key membership)
type stringSet map[interface{}]struct{}

// NewStringSet returns a new instance of a string set implementing the Set interface
func NewStringSet() Set {
	return &stringSet{}
}

// Add the string to the set (returns false if already exists)
func (set *stringSet) Add(i interface{}) bool {
	_, exists := (*set)[i]
	if exists {
		return false
	}

	// create a key with an "empty value" to indicate membership
	(*set)[i] = struct{}{}
	return true
}

// Contains checks to see that all the args are contained in the
// set; it returns false if any arg is not contained
// NOTE: this is not necessary to pass the tests for this solution and
// is only presented here as an example of a minimal set implementation
func (set *stringSet) Contains(i ...interface{}) bool {
	for _, val := range i {
		if _, exists := (*set)[val]; !exists {
			return false
		}
	}

	return true
}

// Remove deletes the value from the set and returns true; else
// returns false if the value was not a member of the set
// NOTE: this is not necessary to pass the tests for this solution and
// is only presented here as an example of a minimal set implementation
func (set *stringSet) Remove(i interface{}) bool {
	if _, exists := (*set)[i]; exists {
		delete(*set, i)
		return true
	}
	return false
}
