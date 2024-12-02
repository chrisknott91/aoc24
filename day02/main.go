package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("inputs/day02.txt")
	if err != nil {
		fmt.Println("error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	safeCount := 0

	for scanner.Scan() {
		line := scanner.Text()
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
