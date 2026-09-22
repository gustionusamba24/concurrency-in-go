package main

import (
	"sync"
	"testing"
)

func TestDataRaceCondition(t *testing.T) {
	var state int32
	var mu sync.RWMutex

	for i := 0; i < 10; i++ {
		go func(i int) {
			mu.Lock()
			state += int32(i)
			// the logic
			mu.Unlock()
		}(i)
	}
}
