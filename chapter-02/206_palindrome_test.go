package chapter02

import (
	"testing"
)

func TestPalindrome(t *testing.T) {
	type sample struct {
		Input    *Node
		Expected bool
	}

	tests := []sample{
		sample{
			Input:    &Node{1, &Node{2, &Node{1, nil}}},
			Expected: true,
		},
		sample{
			Input:    &Node{1, &Node{2, &Node{2, &Node{1, nil}}}},
			Expected: true,
		},
		sample{
			Input:    &Node{1, nil},
			Expected: true,
		},
		sample{
			Input:    &Node{1, &Node{2, &Node{2, &Node{2, nil}}}},
			Expected: false,
		},
	}

	for _, samp := range tests {
		actual := isPalindrome(samp.Input)
		expected := samp.Expected

		if actual != expected {
			t.Errorf("\nexpected %t for %v", expected, samp.Input)
		}
	}
}

func isPalindrome(n *Node) bool {
	if n == nil {
		return false
	}

	rev, length := reverseClone(n)

	for i := 0; i < length/2; i++ {
		if n.Data != rev.Data {
			return false
		}
		n = n.Next
		rev = rev.Next
	}

	return true
}

func reverseClone(n *Node) (*Node, int) {
	var f func(*Node) (*Node, *Node, int)
	f = func(n *Node) (*Node, *Node, int) {
		if n == nil {
			return nil, nil, 0
		}

		clone := &Node{Data: n.Data}
		head, node, count := f(n.Next)
		if node == nil {
			node = clone
			head = node
		} else {
			node.Next = clone
		}
		return head, clone, count + 1
	}

	// we needed to keep track of the current clone node, during recursion
	// but all we really need to return is the new head and the length
	head, _, length := f(n)
	return head, length
}
