package day01

import (
	"advent2025/utils"
	"fmt"
	"math"
)

func Process() {
	dial := 50
	part1 := 0
	part2 := 0

	fileLines := utils.ReadFile("day01/input.txt")

	for _, line := range fileLines {
		value := utils.StringToInt(line[1:])
		if line[0] == 'L' {
			value *= -1
		}
		dialNew := dial + value
		hI := float64(dial) / 100.0
		hF := float64(dialNew) / 100.0
		if value > 0 {
			hI = math.Floor(hI)
			hF = math.Floor(hF)
		} else {
			hI = math.Ceil(hI)
			hF = math.Ceil(hF)
		}
		part2 += utils.Abs(int(hF) - int(hI))
		dial = dialNew

		if dial%100 == 0 {
			part1++
		}
	}

	fmt.Println("Day 1 Results")
	fmt.Println("Part1", part1)
	fmt.Println("Part2", part2)
}
