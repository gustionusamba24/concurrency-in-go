package rw_mutex

import (
	"fmt"
	"sync"
	"testing"

	"github.com/gustionusamba24/concurrency-in-go/internal/format"
)

func Test(t *testing.T) {
	type testCase struct {
		email string
		count int
	}

	runCases := []testCase{
		{"norman@bates.com", 23},
		{"marion@bates.com", 67},
		{"lila@bates.com", 31},
		{"sam@bates.com", 453},
	}

	testCases := runCases

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		sc := safeCounter{
			counts: make(map[string]int),
			mu:     &sync.RWMutex{},
		}
		var wg sync.WaitGroup
		for i := 0; i < test.count; i++ {
			wg.Add(1)
			go func(email string) {
				sc.inc(email)
				wg.Done()
			}(test.email)
		}
		wg.Wait()

		sc.mu.RLock()
		defer sc.mu.RUnlock()
		if output := sc.val(test.email); output != test.count {
			failCount++
			t.Errorf(`---------------------------------
Test Failed:
  email: %v
  count: %v
  expected count: %v
  actual count:   %v
`, test.email, test.count, test.count, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed:
  email: %v
  count: %v
  expected count: %v
  actual count:   %v
`, test.email, test.count, test.count, output)
		}
	}

	format.Summary(len(testCases), passCount, failCount)
}
