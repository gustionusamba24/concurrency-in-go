package main

import "fmt"

func main() {
	myChannel := make(chan string)
	anotherChannel := make(chan string)

	go func() {
		myChannel <- "data"
	}()

	go func() {
		anotherChannel <- "another data"
	}()

	// If the channels both are ready, select will choose one of them at random.
	select {
	case msgFromMyCh := <-myChannel:
		fmt.Println(msgFromMyCh)
	case msgFromAnotherCh := <-anotherChannel:
		fmt.Println(msgFromAnotherCh)
	}
}
