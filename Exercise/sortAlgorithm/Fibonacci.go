package sortalgorithm

func Fibonacci(n int) int {
	if n <= 2 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
