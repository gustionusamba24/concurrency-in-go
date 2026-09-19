package main

import "fmt"

func sliceToChannel(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, num := range nums {
			out <- num
		}
		close(out)
	}()

	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()

	return out
}

func main() {
	// input
	nums := []int{3, 5, 6, 8, 9}

	// stage 1
	dataChannel := sliceToChannel(nums)

	// stage 2
	finalChannel := sq(dataChannel)

	// stage 3
	for n := range finalChannel {
		fmt.Println(n)
	}
}
