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
}

// kthToLastRecursive returns kth node and length of list
// for k=0, kth node is nil; for k=1, kth node is last item in list;
// k=2, second to last item; etc.
func kthToLastRecursive(n *Node, k int) (kth *Node, len int) {
	// base case
	if n == nil {
		return nil, 0
	}

	// keep returning current node until length reaches k
	kth, len = kthToLastRecursive(n.Next, k)
	if len < k {
		return n, len + 1
	}

	// don't update kth anymore now that the kth node has been identified
	// no real need to continue incrementing length, but might as well return
	// something useful (ie, full length of the list traversed)
	return kth, len + 1
}
