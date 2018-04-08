package chapter02

import (
	"testing"
)

func TestRemoveDups(t *testing.T) {
	type sample struct {
		Input    *Node
		Expected *Node
	}

	tests := []sample{
		sample{
			&Node{1, &Node{2, &Node{3, nil}}},
			&Node{1, &Node{2, &Node{3, nil}}},
		},
		sample{
			&Node{1, &Node{2, &Node{3, &Node{3, nil}}}},
			&Node{1, &Node{2, &Node{3, nil}}},
		},
	}

	t.Run("removeDups", func(t *testing.T) {
		for _, samp := range tests {
			// we want to be able to reuse the sample data, so clone it
			actual := samp.Input.Clone()
			expected := samp.Expected

			if ok := removeDups(actual); !ok {
				t.Errorf("problem with input: %v", actual)
			}

			if !ListEqual(actual, expected) {
				t.Errorf("\nactual:   %v\nexpected: %v", actual, expected)
			}
		}
	})

	t.Run("removeDups2", func(t *testing.T) {
		for _, samp := range tests {
			// we want to be able to reuse the sample data, so clone it
			actual := samp.Input.Clone()
			expected := samp.Expected

			if ok := removeDups2(actual); !ok {
				t.Errorf("problem with input: %v", actual)
			}

			if !ListEqual(actual, expected) {
				t.Errorf("\nactual:   %v\nexpected: %v", actual, expected)
			}
		}
	})
}

func removeDups(n *Node) bool {
	set := NewIntSet()

	node := n
	set.Add(node.Data)
	next := node.Next
	for next != nil {
		if set.Add(next.Data) {
			node = next
		}
		next = next.Next
	}
	node.Next = nil
	return true
}

func removeDups2(n *Node) bool {
	for node := n; node != nil; {
		runner := node.Next
		for runner != nil {
			if runner.Data != node.Data {
				break
			} else {
				runner = runner.Next
			}
		}
		node.Next = runner
		node = node.Next
	}

	return true
}
