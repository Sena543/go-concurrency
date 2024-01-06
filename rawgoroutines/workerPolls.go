package rawgoroutines

func Worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- Fib(n)
	}
}

func Fib(n int) int {
	if n <= 1 {
		return n
	}

	return Fib(n-1) + Fib(n-2)
}
