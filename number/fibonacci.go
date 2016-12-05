package number

func TopDownFibonacci(n int) int {
	fibonacci := make([]int, n+1)
	return topDownFibonacci(n, fibonacci)
}

func topDownFibonacci(n int, fibonacci []int) int {
	if fibonacci[n] > 0 {
		return fibonacci[n]
	} else {
		if n == 0 || n == 1 {
			fibonacci[n] = n
		} else {
			fibonacci[n] = topDownFibonacci(n-1, fibonacci) + topDownFibonacci(n-2, fibonacci)
		}
		return fibonacci[n]
	}
}

func BottomUpFibonacci(n int) int {
	fibonacci := make([]int, n+1)
	if n >= 1 {
		fibonacci[0], fibonacci[1] = 0, 1
	}
	for i := 2; i <= n; i++ {
		fibonacci[i] = fibonacci[i-1] + fibonacci[i-2]
	}
	return fibonacci[n]
}
