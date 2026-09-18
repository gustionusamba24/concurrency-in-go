package range_go

import (
	"fmt"
	"slices"
	"testing"

	"github.com/gustionusamba24/concurrency-in-go/internal/format"
)

func TestConcurrentFib(t *testing.T) {
	type testCase struct {
		n        int
		expected []int
	}

	runCases := []testCase{
		{5, []int{0, 1, 1, 2, 3}},
		{3, []int{0, 1, 1}},
		{0, []int{}},
		{1, []int{0}},
		{7, []int{0, 1, 1, 2, 3, 5, 8}},
	}

	testCases := runCases

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		actual := concurrentFib(test.n)
		if !slices.Equal(actual, test.expected) {
			failCount++
			t.Errorf(`---------------------------------
Test Failed:
  n:        %v
  expected: %v
  actual:   %v
`, test.n, test.expected, actual)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed:
  n:        %v
  expected: %v
  actual:   %v
`, test.n, test.expected, actual)
		}
	}

	format.Summary(len(testCases), passCount, failCount)
}
