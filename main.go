package main

import (
	"advent2025/day01"
	"advent2025/day02"
	"advent2025/day03"
	"advent2025/day04"
	"fmt"
	"time"
)

func main() {
	days := []func(){day01.Process, day02.Process, day03.Process, day04.Process}
	totalStart := time.Now()

	for _, function := range days {
		fmt.Println("===================")
		start := time.Now()
		function()
		fmt.Printf("Solved in: %v\n\n", time.Since(start))
	}
	fmt.Println()
	fmt.Printf("Total time: %v\n\n", time.Since(totalStart))
}
