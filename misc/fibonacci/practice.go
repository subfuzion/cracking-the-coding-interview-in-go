package fibonacci

func fib(n int) int {
	if n < 2 {
		return n
	}

	a := 0
	b := 1
	c := 0
	for i := 2; i <= n; i++ {
		c = a + b
		a = b
		b = c
	}
	return c
}

func fibr(n int) int {
	if n < 2 {
		return n
	}
	return fibr(n-2) + fibr(n-1)
}

func fibr_memoized(n int) int {
	if n < 2 {
		return n
	}

	var memo = make([]int, n+1)
	memo[1] = 1

	var f func(int) int
	f = func(n int) int {
		if n < 2 {
			return n
		}

		if memo[n] == 0 {
			memo[n] = f(n-2) + f(n-1)
		}

		return memo[n]
	}

	return f(n-2) + f(n-1)
}

func fibr_memoized2(n int) int {
	if n < 2 {
		return n
	}

	var memo = make([]int, n+1)
	memo[1] = 1

	var f func(int) int
	f = func(n int) int {
		if n < 2 {
			return n
		}

		a := memo[n-2]
		if a == 0 {
			a = f(n - 2)
		}

		b := memo[n-1]
		if b == 0 {
			b = f(n - 1)
		}

		memo[n] = a + b
		return memo[n]
	}

	return f(n-2) + f(n-1)
}
