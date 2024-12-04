package main

import (
	"fmt"
	"github.com/chrisknott91/aoc24/utils"
	"strconv"
	"strings"
)

func main() {
	input := utils.ReadInput("inputs/day02.txt")
	safeCount := 0

	for _, line := range input {
		levels := parseLine(line)
		if isSafe(levels) {
			safeCount++
		} else {
			for i := 0; i < len(levels); i++ {
				levelsCopy := make([]int, len(levels)-1)
				copy(levelsCopy[:i], levels[:i])
				copy(levelsCopy[i:], levels[i+1:])
				if isSafe(levelsCopy) {
					safeCount++
					break
				}
			}
		}
	}

	fmt.Println("Number of safe reports:", safeCount)
}

func parseLine(line string) []int {
	parts := strings.Fields(line)
	levels := make([]int, len(parts))
	for i, part := range parts {
		levels[i], _ = strconv.Atoi(part)
	}
	return levels
}

func isSafe(levels []int) bool {
	if len(levels) < 2 {
		return true
	}

	increasing := levels[0] < levels[1]
	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]
		if diff == 0 || diff < -3 || diff > 3 {
			return false
		}
		if (increasing && levels[i] < levels[i-1]) || (!increasing && levels[i] > levels[i-1]) {
			return false
		}
	}
	return true
}
