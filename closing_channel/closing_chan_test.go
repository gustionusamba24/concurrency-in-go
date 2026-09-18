package closing_channel

import (
	"fmt"
	"testing"

	"github.com/gustionusamba24/concurrency-in-go/internal/format"
)

func TestCountReports(t *testing.T) {
	type testCase struct {
		numBatches int
		expected   int
	}

	runCases := []testCase{
		{3, 114},
		{4, 198},
		{0, 0},
		{1, 15},
		{6, 435},
	}

	testCases := runCases

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		numSentCh := make(chan int)
		go sendReports(test.numBatches, numSentCh)
		output := countReports(numSentCh)
		if output != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Test Failed:
  numBatches: %v
  expected:   %v
  actual:     %v
`, test.numBatches, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed:
  numBatches: %v
  expected:   %v
  actual:     %v
`, test.numBatches, test.expected, output)
		}
	}

	format.Summary(len(testCases), passCount, failCount)
}
