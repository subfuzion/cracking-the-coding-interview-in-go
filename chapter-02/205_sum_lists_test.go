package chapter02

import (
	"testing"
)

func TestSumLists(t *testing.T) {
	type sample struct {
		A             *Node
		B             *Node
		ExpectReverse *Node
		ExpectForward *Node
	}

	tests := []sample{
		sample{
			A:             &Node{7, &Node{1, &Node{6, nil}}},
			B:             &Node{5, &Node{9, &Node{2, nil}}},
			ExpectReverse: &Node{2, &Node{1, &Node{9, nil}}},
			ExpectForward: &Node{1, &Node{3, &Node{0, &Node{8, nil}}}},
		},
		sample{
			A:             &Node{1, &Node{2, nil}},
			B:             &Node{3, nil},
			ExpectReverse: &Node{4, &Node{2, nil}},
			ExpectForward: &Node{1, &Node{5, nil}},
		},
		sample{
			A:             &Node{9, nil},
			B:             &Node{9, nil},
			ExpectReverse: &Node{8, &Node{1, nil}},
			ExpectForward: &Node{1, &Node{8, nil}},
		},
	}

	t.Run("sum_reverse", func(t *testing.T) {
		for _, samp := range tests {
			actual := sumListsReverse(samp.A, samp.B)
			expected := samp.ExpectReverse
			if !actual.Equal(expected) {
				t.Errorf("\nA: %v\nB: %v\nexpected: %v\nactual:   %v", samp.A, samp.B, expected, actual)
			}
		}
	})

	t.Run("sum_forward", func(t *testing.T) {
		for _, samp := range tests {
			actual := sumListsForward(samp.A, samp.B)
			expected := samp.ExpectForward
			if !actual.Equal(expected) {
				t.Errorf("\nA: %v\nB: %v\nexpected: %v\nactual:   %v", samp.A, samp.B, expected, actual)
			}
		}
	})
}

func sumListsReverse(a, b *Node) *Node {
	if a == nil || b == nil {
		return nil
	}

	carry := 0
	var head *Node
	var tail *Node

	pa := a
	pb := b

	for pa != nil && pb != nil {
		sum := pa.Data + pb.Data + carry
		if sum >= 10 {
			carry = sum / 10
			sum = sum % 10
		}
		n := &Node{sum, nil}
		if head == nil {
			head = n
			tail = n
		} else {
			tail.Next = n
			tail = n
		}
		pa = pa.Next
		pb = pb.Next
	}

	var pr *Node
	if pa != nil {
		pr = pa
	} else if pb != nil {
		pr = pb
	}
	for pr != nil {
		sum := pr.Data + carry
		if sum >= 10 {
			carry = sum / 10
			sum = sum % 10
		}
		n := &Node{sum, nil}
		tail.Next = n
		tail = n
		pr = pr.Next
	}

	if carry > 0 {
		n := &Node{carry, nil}
		tail.Next = n
		tail = n
	}

	return head
}

func sumListsForward(a, b *Node) *Node {
	if a == nil || b == nil {
		return nil
	}

	la := a.Length()
	lb := b.Length()
	if la != lb {
		var p **Node
		var nPadding int
		if la < lb {
			p = &a
			nPadding = lb - la
		} else {
			p = &b
			nPadding = la - lb
		}
		for i := 0; i < nPadding; i++ {
			n := &Node{0, *p}
			*p = n
		}
	}

	var f func(a, b *Node) (*Node, int)
	f = func(a, b *Node) (*Node, int) {
		if a == nil {
			return nil, 0
		}

		next, carry := f(a.Next, b.Next)

		sum := a.Data + b.Data + carry
		if sum > 10 {
			carry = sum / 10
			sum = sum % 10
		}
		n := &Node{sum, next}
		return n, carry
	}

	c, carry := f(a, b)
	if carry > 0 {
		n := &Node{carry, c}
		c = n
	}
	return c
}
