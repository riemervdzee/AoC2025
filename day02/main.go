package day02

import (
	"advent2025/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Process() {
	part1 := 0
	part2 := 0
	inputLines := ReadInput("day02/input.txt")
	for _, line := range inputLines {
		parts := strings.Split(line, "-")
		from := utils.StringToInt(parts[0])
		to := utils.StringToInt(parts[1])

		for i := from; i <= to; i++ {
			id := strconv.Itoa(i)
			if checkInvalid(id) {
				val := utils.StringToInt(id)
				part1 += val
				part2 += val
			} else if checkInvalidPart2(id) {
				part2 += utils.StringToInt(id)
			}
		}
	}

	fmt.Println("Day 2 Results")
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func checkInvalid(id string) bool {
	if len(id)%2 == 0 {
		halfIdx := len(id) / 2
		p1 := id[0:halfIdx]
		p2 := id[halfIdx:]
		return p1 == p2
	}
	return false
}

func checkInvalidPart2(id string) bool {
	for i := 0; i < len(id)/2; i++ {
		invalid := true
		if len(id)%(i+1) != 0 {
			continue
		}
		for j := i + 1; j < len(id); j = j + i + 1 {
			a := id[j-i-1 : j]
			b := id[j : j+i+1]
			if a != b {
				invalid = false
				break
			}
		}
		if invalid {
			return true
		}
	}
	return false
}

func ReadInput(filename string) []string {
	file, err := os.Open(filename)
	utils.Check(err)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	scanner.Split(splitAt)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func splitAt(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Return nothing if at EOF
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	// Find the index of the input of comma
	if i := strings.Index(string(data), ","); i >= 0 {
		return i + 1, data[0:i], nil
	}
	// If at EOF + data, return the data
	if atEOF {
		return len(data), data, nil
	}
	return
}
