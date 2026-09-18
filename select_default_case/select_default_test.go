package select_default_case

import (
	"fmt"
	"slices"
	"testing"
	"time"

	"github.com/gustionusamba24/concurrency-in-go/internal/format"
)

func TestSaveBackups(t *testing.T) {
	type testCase struct {
		expectedLogs []string
	}

	runCases := []testCase{
		{
			expectedLogs: []string{
				"Nothing to do, waiting...",
				"Nothing to do, waiting...",
				"Nothing to do, waiting...",
				"Taking a backup snapshot...",
				"Nothing to do, waiting...",
				"Nothing to do, waiting...",
				"Nothing to do, waiting...",
				"Taking a backup snapshot...",
				"Nothing to do, waiting...",
				"All backups saved!",
			},
		},
	}

	testCases := runCases

	passed, failed := 0, 0
	for _, test := range testCases {
		expectedLogs := test.expectedLogs

		snapshotTicker := time.Tick(1500 * time.Millisecond)
		saveAfter := time.After(3500 * time.Millisecond)
		logChan := make(chan string)
		go saveBackups(snapshotTicker, saveAfter, logChan)
		actualLogs := []string{}
		for actualLog := range logChan {
			fmt.Println(actualLog)
			actualLogs = append(actualLogs, actualLog)
		}

		if !slices.Equal(expectedLogs, actualLogs) {
			t.Errorf(`---------------------------------
Test Failed:
expected:
%v
actual:
%v
`, sliceWithBullets(expectedLogs), sliceWithBullets(actualLogs))
			failed++
		} else {
			fmt.Printf(`---------------------------------
Test Passed:
expected:
%v
actual:
%v
`, sliceWithBullets(expectedLogs), sliceWithBullets(actualLogs))
			passed++
		}
	}

	format.Summary(len(testCases), passed, failed)
}

func sliceWithBullets[T any](slice []T) string {
	if slice == nil {
		return "  <nil>"
	}
	if len(slice) == 0 {
		return "  []"
	}
	output := ""
	for i, item := range slice {
		form := "  - %#v\n"
		if i == (len(slice) - 1) {
			form = "  - %#v"
		}
		output += fmt.Sprintf(form, item)
	}
	return output
}
