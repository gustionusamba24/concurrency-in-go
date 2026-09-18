package range_go

func concurrentFib(n int) []int {
	ch := make(chan int)
	go fibonacci(n, ch)
	var result []int
	for num := range ch {
		result = append(result, num)
	}
	return result
}

// don't touch below the line
func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}
