package main

import (
	"fmt"
	"github.com/chrisknott91/aoc24/utils"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := utils.ReadInput("inputs/day01.txt")
	left, right := splitInput(input)

	fmt.Println("Total Distance:", totalDistance(left, right))
	fmt.Println("Similarity Score:", similarityScore(left, right))
}
func splitInput(input []string) (left []int, right []int) {
	for _, line := range input {
		parts := strings.Fields(line)
		if len(parts) != 2 {
			log.Fatalf("invalid input line: %v", line)
		}
		l, err1 := strconv.Atoi(parts[0])
		r, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			log.Fatalf("error parsing integers: %v, %v", err1, err2)
		}
		left = append(left, l)
		right = append(right, r)
	}

	return left, right
}

func similarityScore(left, right []int) int {
	rightCount := make(map[int]int)
	for _, num := range right {
		rightCount[num]++
	}

	similarityScore := 0
	for _, num := range left {
		similarityScore += num * rightCount[num]
	}
	return similarityScore
}

func totalDistance(left, right []int) int {
	sort.Ints(left)
	sort.Ints(right)

	totalDistance := 0
	for i := range left {
		totalDistance += int(math.Abs(float64(left[i] - right[i])))
	}
	return totalDistance
}
