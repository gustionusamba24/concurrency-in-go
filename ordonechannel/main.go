package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	done := make(chan interface{})
	defer close(done)

	cows := make(chan interface{}, 100)
	wolves := make(chan interface{}, 100)

	go func() {
		for {
			select {
			case <-done:
				return
			case cows <- "mooo":
			}
		}
	}()

	go func() {
		for {
			select {
			case <-done:
				return
			case wolves <- "woof":
			}
		}
	}()

	wg.Add(1)
	go consumeCows(done, cows)

	wg.Add(1)
	go consumeWolves(done, wolves)

	wg.Wait()
}

func consumeCows(done <-chan interface{}, cows <-chan interface{}) {
	defer wg.Done()
	for cow := range orDone(done, cows) {
		// do some complex logic
		fmt.Println(cow)
	}
}

func consumeWolves(done <-chan interface{}, wolves <-chan interface{}) {
	defer wg.Done()
	for wolf := range orDone(done, wolves) {
		// do some complex logic
		fmt.Println(wolf)
	}
}

func orDone(done, c <-chan interface{}) <-chan interface{} {
	relayStream := make(chan interface{})
	go func() {
		defer close(relayStream)
		for {
			select {
			case <-done:
				return
			case v, ok := <-c:
				if !ok {
					return
				}
				select {
				case relayStream <- v:
				case <-done:
					return
				}
			}
		}
	}()

	return relayStream
}
