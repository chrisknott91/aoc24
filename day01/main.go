package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

func readInput(filePath string) (left []int, right []int) {
	file, fileError := os.Open(filePath)
	if fileError != nil {
		log.Fatalf("error opening file: %v", fileError)
	}
	defer file.Close()

	var l, r int
	var err error
	for err == nil {
		if _, err = fmt.Fscanf(file, "%d %d", &l, &r); err == nil {
			left = append(left, l)
			right = append(right, r)
		}
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

func main() {
	left, right := readInput("inputs/day01.txt")

	fmt.Println("Total Distance:", totalDistance(left, right))
	fmt.Println("Similarity Score:", similarityScore(left, right))
}
