package format

import "fmt"

func Summary(totalTests, passCount, failCount int) {
	fmt.Println("===============================")
	fmt.Println("TEST SUMMARY")
	fmt.Printf("Total Tests: %d\n", totalTests)
	fmt.Printf("Passed: %d\n", passCount)
	fmt.Printf("Failed: %d\n", failCount)
	fmt.Println("===============================")
}
