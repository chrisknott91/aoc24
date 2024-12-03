package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("inputs/day03.txt")
	if err != nil {
		fmt.Println("error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var multiplierPairs [][]int
	var mulEnabled bool = true

	for scanner.Scan() {
		line := scanner.Text()
		pairs, mulEnabledStatus := parseLine(line, mulEnabled)
		mulEnabled = mulEnabledStatus
		multiplierPairs = append(multiplierPairs, pairs...)
	}

	result := 0
	for _, pair := range multiplierPairs {
		result += pair[0] * pair[1]
	}

	fmt.Println("Pairs", multiplierPairs)
	fmt.Println("Result of multiplications:", result)
}

func parseLine(line string, mulEnabled bool) ([][]int, bool) {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
	matches := re.FindAllStringSubmatch(line, -1)

	var result [][]int

	for _, match := range matches {
		if match[0] == "do()" {
			mulEnabled = true
		} else if match[0] == "don't()" {
			mulEnabled = false
		}

		if mulEnabled {
			x, _ := strconv.Atoi(match[1])
			y, _ := strconv.Atoi(match[2])
			result = append(result, []int{x, y})
		}
	}
	return result, mulEnabled
}
