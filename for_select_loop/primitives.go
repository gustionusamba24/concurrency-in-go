package main

import "fmt"

func main() {
	buffChannel := make(chan string, 3)
	names := []string{"Alice", "Bob", "Charlie"}

	for _, name := range names {
		select {
		case buffChannel <- name:
		}
	}

	close(buffChannel)

	for result := range buffChannel {
		fmt.Println(result)
	}
}
