package day03

import (
	"advent2025/utils"
	"fmt"
)

func Process() {
	part1 := 0
	part2 := 0

	grid := utils.ReadFileAsGrid("day03/input.txt")
	for _, row := range grid {
		part1 += joltBatteries(0, row, 2)
		part2 += joltBatteries(0, row, 12)
	}

	fmt.Println("Day 3 Results")
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func joltBatteries(accumulated int, batteries []uint8, count int) int {
	batterySlice := batteries[0 : len(batteries)-count+1]
	maxIndex, maxValue := findLargestBattery(batterySlice)
	accumulated = accumulated*10 + int(maxValue) - '0'
	count--
	if count == 0 {
		return accumulated
	}

	return joltBatteries(accumulated, batteries[maxIndex+1:], count)
}

func findLargestBattery(batteries []uint8) (int, uint8) {
	maxValue := uint8(0)
	maxIndex := 0
	for i := 0; i < len(batteries); i++ {
		if batteries[i] > maxValue {
			maxValue = batteries[i]
			maxIndex = i
		}
	}
	return maxIndex, maxValue
}
