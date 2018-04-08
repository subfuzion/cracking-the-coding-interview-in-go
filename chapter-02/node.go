package chapter02

import (
	"strconv"
	"strings"
)

// Node for single linked list of integers
type Node struct {
	Data int
	Next *Node
}

// Length returns length of single linked list starting from n
func (n *Node) Length() int {
	count := 0
	for node := n; node != nil; {
		if node != nil {
			count++
		}
		node = node.Next
	}
	return count
}

// Equal returns true if n.Data == n2.Data
func (n *Node) Equal(n2 *Node) bool {
	switch {
	case n == n2:
		return true
	case n != nil && n2 != nil:
		return n.Data == n2.Data
	}
	return false
}

// String stringifies the lists starting from node
func (n Node) String() string {
	b := strings.Builder{}
	b.WriteRune('{')
	next := &n
	for next != nil {
		b.WriteString(strconv.Itoa(next.Data))
		if next.Next != nil {
			b.WriteString(", ")
		}
		next = next.Next
	}
	b.WriteRune('}')
	return b.String()
}

// Clone returns a copy of the list with new Node instances
func (n Node) Clone() *Node {
	head := &Node{Data: n.Data}

	for clone, node := head, n.Next; node != nil; {
		clone.Next = &Node{Data: node.Data, Next: node.Next}
		clone = clone.Next
		node = node.Next
	}

	return head
}

// ListEqual iterates lists to check if Data fields are equal
// returns false if one list is shorter than the other
func ListEqual(n1 *Node, n2 *Node) bool {
	if n1 == nil || n2 == nil {
		return false
	}

	p1 := n1
	p2 := n2
	for p1.Data == p2.Data {
		p1 = p1.Next
		p2 = p2.Next

		// if either list ends before the other, then lists aren't equal
		if (p1 == nil && p2 != nil) || (p1 != nil && p2 == nil) {
			return false
		}

		// nothing left, so lists are equal
		if p1 == nil && p2 == nil {
			return true
		}
	}

	return false
}
