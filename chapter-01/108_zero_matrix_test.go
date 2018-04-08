package chapter01

import (
	"testing"
)

// matrix is already defined in rotate_matrix_test.go
// type matrix [][]int

//equal is already defined in rotate_matrix_test.go
// func (m matrix) equal(m2 matrix) bool {
// 	if len(m) != len(m2) || len(m[0]) != len(m2[0]) {
// 		return false
// 	}

// 	for i, row := range m {
// 		for j, val := range m2 {
// 			if val != m2[i][j] {
// 				return false
// 			}
// 		}
// 	}

// 	return true
// }

func TestZeroMatrix(t *testing.T) {
	type sample struct {
		Input    matrix
		Expected matrix
	}

	tests := []sample{
		sample{
			Input: matrix{
				{1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1},
			},
			Expected: matrix{
				{1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1},
			},
		},
		sample{
			Input: matrix{
				{1, 1, 1, 1, 1, 1},
				{1, 1, 0, 1, 1, 1},
				{1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 0, 1},
				{1, 1, 1, 1, 1, 1},
			},
			Expected: matrix{
				{1, 1, 0, 1, 0, 1},
				{0, 0, 0, 0, 0, 0},
				{1, 1, 0, 1, 0, 1},
				{0, 0, 0, 0, 0, 0},
				{1, 1, 0, 1, 0, 1},
			},
		},
		sample{
			Input: matrix{
				{1, 0, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1},
				{0, 1, 1, 1, 1, 1},
			},
			Expected: matrix{
				{0, 0, 0, 0, 0, 0},
				{0, 0, 1, 1, 1, 1},
				{0, 0, 1, 1, 1, 1},
				{0, 0, 1, 1, 1, 1},
				{0, 0, 0, 0, 0, 0},
			},
		},
		sample{
			Input: matrix{
				{0, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1},
			},
			Expected: matrix{
				{0, 0, 0, 0, 0, 0},
				{0, 1, 1, 1, 1, 1},
				{0, 1, 1, 1, 1, 1},
				{0, 1, 1, 1, 1, 1},
				{0, 1, 1, 1, 1, 1},
			},
		},
	}

	for _, samp := range tests {
		actual := samp.Input
		expected := samp.Expected

		if ok := nullifyMatrix(actual); !ok {
			t.Errorf("can't nullify matrix (check data)\n%v", actual)
		}

		if !actual.equal(expected) {
			t.Errorf("actual != expected:\nactual:\n%vexpected:\n%v", actual, expected)
		}
	}
}

func nullifyMatrix(m matrix) bool {
	if len(m) == 0 || len(m[0]) == 0 {
		return false
	}

	// we will use first row and first column for storage, so first
	// flag whether or not first row or first column has any zeroes
	nullifyFirstColumn := false
	for i := range m {
		if m[i][0] == 0 {
			nullifyFirstColumn = true
		}
	}
	nullifyFirstRow := false
	for j := range m[0] {
		if m[0][j] == 0 {
			nullifyFirstRow = true
		}
	}

	// // flag rows and columns to nullify
	for i := 1; i < len(m); i++ {
		for j := 1; j < len(m[0]); j++ {
			if m[i][j] == 0 {
				m[i][0] = 0
				m[0][j] = 0
			}
		}
	}

	// nullify flagged rows
	for i := 1; i < len(m); i++ {
		if m[i][0] == 0 {
			for j := 1; j < len(m[0]); j++ {
				m[i][j] = 0
			}
		}
	}

	// // nullify flagged columns
	for j := 1; j < len(m[0]); j++ {
		if m[0][j] == 0 {
			for i := 1; i < len(m); i++ {
				m[i][j] = 0
			}
		}
	}

	if nullifyFirstColumn {
		for i := range m {
			m[i][0] = 0
		}
	}

	if nullifyFirstRow {
		for j := range m[0] {
			m[0][j] = 0
		}
	}

	return true
}
