package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		for {
			select {
			default:
				fmt.Println("DO A LOT OF WORK")
			}
		}
	}()

	time.Sleep(time.Second * 10)
}
