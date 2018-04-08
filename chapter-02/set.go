package chapter02

// IntSet is a set of integers
type IntSet map[int]struct{}

// NewIntSet returns a new (empty) set for integers
func NewIntSet() *IntSet {
	return &IntSet{}
}

// Add n to the set and return true; if already present in
// the set then return false
func (set *IntSet) Add(n int) bool {
	if _, exists := (*set)[n]; exists {
		return false
	}
	(*set)[n] = struct{}{}
	return true
}
