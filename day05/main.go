package main

import (
	"fmt"
	"github.com/chrisknott91/aoc24/utils"
	"strconv"
	"strings"
)

func main() {
	input := utils.ReadInput("inputs/day05.txt")
	rules, updates := parseInput(input)

	validUpdates := filterUpdates(rules, updates, true)
	middleSumValid := sumMiddlePages(validUpdates)
	fmt.Println("Sum of middle pages (valid updates):", middleSumValid)

	invalidUpdates := filterUpdates(rules, updates, false)
	correctedUpdates := correctUpdates(rules, invalidUpdates)
	middleSumCorrected := sumMiddlePages(correctedUpdates)
	fmt.Println("Sum of middle pages (corrected updates):", middleSumCorrected)
}

func parseInput(input []string) (map[int][]int, [][]int) {
	rules := make(map[int][]int)
	var updates [][]int
	isUpdateSection := false

	for _, line := range input {
		if line == "" {
			isUpdateSection = true
			continue
		}

		if isUpdateSection {
			update := parseUpdate(line)
			updates = append(updates, update)
		} else {
			rule := parseRule(line)
			rules[rule[0]] = append(rules[rule[0]], rule[1])
		}
	}

	return rules, updates
}

func parseRule(line string) [2]int {
	parts := strings.Split(line, "|")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return [2]int{x, y}
}

func parseUpdate(line string) []int {
	parts := strings.Split(line, ",")
	update := make([]int, len(parts))
	for i, part := range parts {
		update[i], _ = strconv.Atoi(part)
	}
	return update
}

func filterUpdates(rules map[int][]int, updates [][]int, isValid bool) [][]int {
	var filteredUpdates [][]int

	for _, update := range updates {
		if isValidUpdate(rules, update) == isValid {
			filteredUpdates = append(filteredUpdates, update)
		}
	}

	return filteredUpdates
}

func isValidUpdate(rules map[int][]int, update []int) bool {
	position := make(map[int]int)
	for i, page := range update {
		position[page] = i
	}

	for x, ys := range rules {
		for _, y := range ys {
			if posX, okX := position[x]; okX {
				if posY, okY := position[y]; okY {
					if posX >= posY {
						return false
					}
				}
			}
		}
	}

	return true
}

func correctUpdates(rules map[int][]int, updates [][]int) [][]int {
	var correctedUpdates [][]int

	for _, update := range updates {
		correctedUpdate := topologicalSort(rules, update)
		correctedUpdates = append(correctedUpdates, correctedUpdate)
	}

	return correctedUpdates
}

func topologicalSort(rules map[int][]int, update []int) []int {
	inDegree := make(map[int]int)
	graph := make(map[int][]int)
	for _, page := range update {
		inDegree[page] = 0
	}

	for x, ys := range rules {
		for _, y := range ys {
			if contains(update, x) && contains(update, y) {
				graph[x] = append(graph[x], y)
				inDegree[y]++
			}
		}
	}

	var queue []int
	for page, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, page)
		}
	}

	var sorted []int
	for len(queue) > 0 {
		page := queue[0]
		queue = queue[1:]
		sorted = append(sorted, page)

		for _, neighbor := range graph[page] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return sorted
}

func contains(slice []int, item int) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func sumMiddlePages(updates [][]int) int {
	sum := 0
	for _, update := range updates {
		middleIndex := len(update) / 2
		sum += update[middleIndex]
	}
	return sum
}
