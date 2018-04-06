package chapter01

import (
	"fmt"
	"strings"
	"testing"
)

type matrix [][]int

func (m matrix) equal(m2 matrix) bool {
	if m2 == nil {
		return false
	}

	if len(m) != len(m2) || len(m[0]) != len(m2[0]) {
		return false
	}

	for i, row := range m {
		for j, val := range row {
			if val != m2[i][j] {
				return false
			}
		}
	}

	return true
}

func (m matrix) square() bool {
	if len(m) != 0 && len(m) == len(m[0]) {
		return true
	}
	return false
}

func (m matrix) String() string {
	b := strings.Builder{}
	for _, row := range m {
		for _, val := range row {
			b.WriteString(fmt.Sprintf("%4d", val))
		}
		b.WriteString("\n")
	}
	return b.String()
}

func TestRotateMatrix(t *testing.T) {
	type sample struct {
		Input    matrix
		Expected matrix
	}

	tests := []sample{
		sample{
			Input: matrix{
				{0, 1, 2},
				{3, 4, 5},
				{6, 7, 8},
			},
			Expected: matrix{
				{6, 3, 0},
				{7, 4, 1},
				{8, 5, 2},
			},
		},
	}

	for _, samp := range tests {
		actual := samp.Input
		expected := samp.Expected

		// if valid input then matrix is rotated in place
		ok := rotate(samp.Input)

		if !ok {
			t.Errorf("error rotating matrix (must be square with non-zero dimensions):\n%v", samp.Input)
		}

		if !actual.equal(expected) {
			t.Errorf("actual != expected\nActual:\n%sExpected:\n%s", actual, expected)
		}
	}
}

func rotate(m matrix) bool {
	if !m.square() {
		return false
	}

	// work our way layer by layer toward the middle
	n := len(m)
	for layer := 0; layer < n; layer++ {
		// this is a square matrix, so first and last identify starting and ending cell
		// for row of each layer
		// Ex: 6x6 matrix
		// iter 0:  first=0 last=5
		// iter 1:  first=1 last=4
		// iter 2:  first=2 last=3
		first := layer
		last := n - 1 - layer

		for i := first; i < last; i++ {
			offset := i - first

			// save before overwriting
			top := m[first][i]

			// move left to top
			m[first][i] = m[last-offset][first]

			// move bottom to left
			m[last-offset][first] = m[last][last-offset]

			// move right to bottom
			m[last][last-offset] = m[i][last]

			// move (saved) top to right
			m[i][last] = top
		}
	}
	return true
}
