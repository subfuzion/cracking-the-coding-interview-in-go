package prime

//
// Solution
//

// Prime1
func Prime1(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n > 1
}

// Prime2 cuts loop in half
func Prime2(n int) bool {
	mid := n / 2
	for i := 2; i <= mid; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n > 1
}
