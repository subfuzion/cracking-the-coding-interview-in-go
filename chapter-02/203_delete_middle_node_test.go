package chapter02

import (
	"testing"
)

func TestDeleteMiddleNode(t *testing.T) {
	type sample struct {
		Input    *Node
		Remove   int
		Expected *Node
		Success  bool
	}

	tests := []sample{
		sample{
			Input:    &Node{1, nil},
			Remove:   0,
			Expected: &Node{1, nil},
			Success:  false,
		},
		sample{
			Input:    &Node{1, &Node{2, nil}},
			Remove:   0,
			Expected: &Node{1, &Node{2, nil}},
			Success:  false,
		},
		sample{
			Input:    &Node{1, &Node{2, nil}},
			Remove:   1,
			Expected: &Node{1, &Node{2, nil}},
			Success:  false,
		},
		sample{
			Input:    &Node{1, &Node{2, &Node{3, nil}}},
			Remove:   1,
			Expected: &Node{1, &Node{3, nil}},
			Success:  true,
		},
		sample{
			Input:    &Node{1, &Node{2, &Node{3, &Node{4, nil}}}},
			Remove:   2,
			Expected: &Node{1, &Node{2, &Node{4, nil}}},
			Success:  true,
		},
		sample{
			Input:    &Node{1, &Node{2, &Node{3, &Node{4, nil}}}},
			Remove:   4,
			Expected: &Node{1, &Node{2, &Node{4, nil}}},
			Success:  false,
		},
	}

	for _, samp := range tests {
		actual := samp.Input
		ok := removeMiddle(actual, samp.Remove)
		if ok != samp.Success {
			t.Errorf("ok != %t for input: %v", samp.Success, actual)
		}

		expected := samp.Expected
		if !actual.Equal(expected) {
			t.Errorf("actual != expected\actual: %v\nexpected: %v", actual, expected)
		}
	}
}

func removeMiddle(node *Node, index int) bool {
	// first or last
	if index == 0 || node == nil || node.Next == nil {
		return false
	}

	for i := 0; i < index; i++ {
		node = node.Next
		// if next is nil, then this the end of the list, so this is not a middle node
		if node.Next == nil {
			return false
		}
	}

	node.Data = node.Next.Data
	node.Next = node.Next.Next
	return true
}
