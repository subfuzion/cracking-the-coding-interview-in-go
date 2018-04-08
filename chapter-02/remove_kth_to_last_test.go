package chapter02

import (
	"testing"
)

func TestKthToLast(t *testing.T) {
	type sample struct {
		Input    *Node
		K        int
		Expected *Node
	}

	tests := []sample{
		sample{
			Input:    &Node{1, nil},
			K:        0,
			Expected: nil,
		},
		sample{
			Input:    &Node{1, nil},
			K:        1,
			Expected: &Node{1, nil},
		},
		sample{
			Input:    &Node{1, nil},
			K:        2,
			Expected: nil,
		},
		sample{
			Input:    &Node{1, &Node{2, &Node{3, &Node{4, nil}}}},
			K:        0,
			Expected: nil,
		},
		sample{
			Input:    &Node{1, &Node{2, &Node{3, &Node{4, nil}}}},
			K:        1,
			Expected: &Node{4, nil},
		},
		sample{
			Input:    &Node{1, &Node{2, &Node{3, &Node{4, nil}}}},
			K:        2,
			Expected: &Node{3, &Node{4, nil}},
		},
	}

	t.Run("recursive", func(t *testing.T) {
		for _, samp := range tests {
			actual, _ := kthToLastRecursive(samp.Input, samp.K)
			expected := samp.Expected

			if actual != nil && actual.Length() != expected.Length() {
				t.Errorf("k=%d\ninput:\n%v\nexpected:\n%v\nactual:\n%v", samp.K, samp.Input, expected, actual)
			}

			if !actual.Equal(expected) {
				t.Errorf("k=%d\ninput:\n%v\nexpected:\n%v\nactual:\n%v", samp.K, samp.Input, expected, actual)
			}
		}
	})

	t.Run("iterative", func(t *testing.T) {
		for _, samp := range tests {
			actual := kthToLastIterative(samp.Input, samp.K)
			expected := samp.Expected

			if !actual.Equal(expected) {
				t.Errorf("k=%d\ninput:\n%v\nexpected:\n%v\nactual:\n%v", samp.K, samp.Input, expected, actual)
			}
		}
	})
}

// kthToLastRecursive returns kth node and length of list
// for k=0, kth node is nil; for k=1, kth node is last item in list;
// k=2, second to last item; etc.
func kthToLastRecursive(n *Node, k int) (*Node, int) {
	length := 0

	f := func(n *Node, k int) (*Node, int) {
		// base case
		if n == nil {
			return nil, 0
		}

		// keep returning current node until length reaches k
		kth, length := kthToLastRecursive(n.Next, k)
		if length < k {
			return n, length + 1
		}

		// don't update kth anymore now that the kth node has been identified
		// no real need to continue incrementing length, but might as well return
		// something useful (ie, full length of the list traversed)
		return kth, length + 1
	}

	// make sure list was actually long enough for kth node
	kth, length := f(n, k)
	if length < k {
		return nil, length
	}
	return kth, length
}

// Non-recursive version
func kthToLastIterative(n *Node, k int) *Node {
	if n == nil || k == 0 {
		return nil
	}

	p2 := n
	for i := 0; i < k; i++ {
		if p2 == nil {
			return nil
		}
		p2 = p2.Next
	}

	for p2 != nil {
		n = n.Next
		p2 = p2.Next
	}

	return n
}
