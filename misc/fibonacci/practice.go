package fibonacci

func fibr(n int) int {
	if n < 2 {
		return n
	}

	return fibr(n-2) + fibr(n-1)
}

func fibr_memo(n int) int {
	if n < 2 {
		return n
	}

	memo := make([]int, n+1)
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

func fib(n int) int {
	if n < 2 {
		return n
	}

	a := 0
	b := 1
	sum := 0
	for i := 2; i <= n; i++ {
		sum = a + b
		a = b
		b = sum
	}
	return sum
}

func fibx(N int) int {
	arr := []int{0, 1}

	for i := 2; i <= N; i++ {
		arr[i%2] = arr[0] + arr[1]
	}

	return arr[N%2]
}
