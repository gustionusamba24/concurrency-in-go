package main

import (
	"fmt"

	"github.com/gustionusamba24/concurrency-in-go/concurrency"
	"github.com/gustionusamba24/concurrency-in-go/select_go"
)

func main() {
	fmt.Println("This repository explain the concept of Concurrency in Go")

	fmt.Println("Intro to Concurrency in Go")
	concurrency.RunProgram()

	fmt.Println("====================END OF SECTION====================")
	fmt.Println()

	fmt.Println("Select in Golang")
	select_go.RunProgram()
}
