package main

import (
	"fmt"
	"math/rand"
)

func repeatFunc[T any, K any](done <-chan K, fn func() T) <-chan T {
	stream := make(chan T)
	go func() {
		defer close(stream)
		for {
			select {
			case <-done:
				return
			case stream <- fn():
			}
		}
	}()

	return stream
}

func take[T any, K any](done <-chan K, stream <-chan T, num int) <-chan T {
	taken := make(chan T)
	go func() {
		defer close(taken)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case taken <- <-stream:
			}
		}
	}()

	return taken
}

func main() {
	done := make(chan int)
	defer close(done)

	randNumFetcher := func() int {
		return rand.Intn(500000000)
	}

	for rando := range take(done, repeatFunc(done, randNumFetcher), 10) {
		fmt.Println(rando)
	}
}
