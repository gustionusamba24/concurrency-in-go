package main

import (
	"fmt"
	"sync"
	"time"
)

func firstMain() {
	start := time.Now()
	userID := 10
	respch := make(chan string, 128)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go fetchUserData(userID, respch, wg)

	wg.Add(1)
	go fetchUserRecommendations(userID, respch, wg)

	wg.Add(1)
	go fetchUserLikes(userID, respch, wg)

	wg.Wait()

	close(respch)

	for resp := range respch {
		fmt.Println(resp)
	}

	fmt.Println(time.Since(start))
}

func fetchUserData(userID int, respch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 80)

	respch <- "user data"
}

func fetchUserRecommendations(userID int, respch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 120)

	respch <- "user recommendations"
}

func fetchUserLikes(userID int, respch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 50)

	respch <- "user likes"
}
