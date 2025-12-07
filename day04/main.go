package day04

import (
	"advent2025/utils"
	"fmt"
)

func Process() {
	part1 := 0
	part2 := 0

	grid := utils.ReadFileAsGrid("day04/input.txt")
	result := MarkForCleanup(grid)
	part1 = result
	part2 = result
	Cleanup(grid)
	for result > 0 {
		result = MarkForCleanup(grid)
		part2 += result
		Cleanup(grid)
	}

	fmt.Println("Day 4 Results")
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func MarkForCleanup(grid utils.Grid) int {
	marked := 0
	for y, row := range grid {
		for x, cell := range row {
			if cell != '@' {
				continue
			}
			hits := 0
			for _, vec := range utils.AllDirections {
				position := utils.VectorAdd(utils.Vector{x, y}, vec)
				if utils.GridValidPosition(grid, position) && grid[position[1]][position[0]] != '.' {
					hits++
				}
			}
			if hits < 4 {
				marked++
				grid[y][x] = 'x'
			}
		}
	}
	return marked
}

func Cleanup(grid utils.Grid) {
	hits := 0
	for y, row := range grid {
		for x, cell := range row {
			if cell == 'x' {
				grid[y][x] = '.'
				hits++
			}
		}
	}
}
