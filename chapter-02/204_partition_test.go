package chapter02

import (
	"testing"
)

func TestPartition(t *testing.T) {
	type sample struct {
		ID       int
		Input    *Node
		Pivot    int
		TestFunc func(*Node, int) bool
	}

	tests := []sample{
		sample{
			ID:       1,
			Input:    &Node{1, &Node{2, &Node{3, nil}}},
			Pivot:    1,
			TestFunc: testFunc,
		},
		sample{
			ID:       2,
			Input:    &Node{1, &Node{2, &Node{3, nil}}},
			Pivot:    2,
			TestFunc: testFunc,
		},
		sample{
			ID:       3,
			Input:    &Node{1, &Node{2, &Node{3, nil}}},
			Pivot:    3,
			TestFunc: testFunc,
		},
		sample{
			ID:       4,
			Input:    &Node{3, &Node{8, &Node{5, &Node{10, &Node{2, &Node{1, nil}}}}}},
			Pivot:    5,
			TestFunc: testFunc,
		},
	}

	for _, samp := range tests {
		actual := partition(samp.Input.Clone(), samp.Pivot)

		// to see log output, use -v
		// $ go test -v -timeout 30s -run ^TestPartition$
		t.Logf("test #%d\nInput    : %v\nPartioned: %v", samp.ID, samp.Input, actual)

		if ok := samp.TestFunc(actual, samp.Pivot); !ok {
			t.Errorf("test #%d failed", samp.ID)
		}

	}
}

func testFunc(n *Node, pivot int) bool {
	pivoted := false

	for next := n; next != nil; {
		if next.Data >= pivot {
			pivoted = true
		}
		if next.Data < pivot && pivoted {
			return false
		}
		next = next.Next
	}

	return true
}

func partition(n *Node, pivot int) *Node {
	if n == nil {
		return nil
	}

	// insert values < pivot at head
	head := n
	// insert values >= pivot at tail
	tail := n

	for node := n; node != nil; {
		next := node.Next
		if node.Data < pivot {
			node.Next = head
			head = node
		} else {
			node.Next = nil
			tail.Next = node
			tail = node
		}
		node = next
	}

	return head
}
