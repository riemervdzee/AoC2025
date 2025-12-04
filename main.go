package main

import (
	"advent2025/day01"
	"fmt"
	"time"
)

func main() {
	days := []func(){day01.Process}
	totalStart := time.Now()

	for _, function := range days {
		function()
	}
	fmt.Println()
	fmt.Printf("Total time: %v\n\n", time.Since(totalStart))
}
