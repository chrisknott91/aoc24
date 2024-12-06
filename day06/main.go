package main

import (
	"fmt"
	"github.com/chrisknott91/aoc24/utils"
)

type Direction int

// Directions
const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)

// main reads the input, calculates the number of distinct positions visited by the guard, and the number of loop-causing obstructions
func main() {
	input := utils.ReadInput("inputs/day06.txt")
	fmt.Println("Number of distinct positions visited: ", calculateVisitedPositions(input))
	fmt.Println("Number of loop-causing obstructions: ", findLoopCausingObstructions(input))
}

// calculateVisitedPositions calculates the number of distinct positions visited by the guard
func calculateVisitedPositions(input []string) int {
	visited := make(map[[2]int]bool)   // Store visited positions
	obstacles := make(map[[2]int]bool) // Store obstacles
	var posX, posY int                 // Starting position
	var direction Direction = UP       // Starting direction

	// Initialize the obstacles and the starting position
	for y, row := range input {
		for x, char := range row {
			// If the character is an obstacle, add it to the obstacles map
			if char == '#' {
				obstacles[[2]int{y, x}] = true
			}
			// If the character is the starting position, set the position
			if char == '^' {
				posY, posX = y, x
				visited[[2]int{y, x}] = true
			}
		}
	}

	// Ensure the starting position is counted as visited
	if _, ok := visited[[2]int{posY, posX}]; !ok {
		visited[[2]int{posY, posX}] = true
	}

	// Simulate the guard's movement
	for {
		nextPos := getNextPosition(posY, posX, direction)
		// If the next position is outside the grid, break the loop
		if isOutsideGrid(nextPos, input) {
			break
		}
		// If the next position is an obstacle, turn right
		if obstacles[nextPos] {
			direction = turnRight(direction)
		} else { // Otherwise, move to the next position
			visited[nextPos] = true
			posY, posX = nextPos[0], nextPos[1]
		}
	}

	// Return the number of visited distinct positions
	return len(visited)
}

// findLoopCausingObstructions finds the number of loop-causing obstructions
func findLoopCausingObstructions(input []string) int {
	var sum int // Sum of the number of positions visited more than once
	var orgPosX, orgPosY int
	origObstacles := make(map[[2]int]bool)

	// Initialize the obstacles and the starting position
	for y, row := range input {
		for x, char := range row {
			if char == '#' {
				origObstacles[[2]int{y, x}] = true
			}
			if char == '^' {
				orgPosY, orgPosX = y, x
			}
		}
	}

	// Check each position in the grid
	for i := range input {
		for j := range input[i] {
			if i == orgPosY && j == orgPosX {
				continue
			}

			visited := make(map[[3]int]bool)
			var direction Direction = UP
			posX, posY := orgPosX, orgPosY
			visited[[3]int{posY, posX, int(direction)}] = true
			obstacles := make(map[[2]int]bool)
			for k, v := range origObstacles {
				obstacles[k] = v
			}
			obstacles[[2]int{i, j}] = true

			var loop, finished bool
			for !finished && !loop {
				nextPos := getNextPosition(posY, posX, direction)
				if obstacles[nextPos] {
					direction = turnRight(direction)
				} else if isOutsideGrid(nextPos, input) {
					finished = true
				} else {
					if visited[[3]int{nextPos[0], nextPos[1], int(direction)}] {
						loop = true
						finished = true
					}
					visited[[3]int{nextPos[0], nextPos[1], int(direction)}] = true
					posY, posX = nextPos[0], nextPos[1]
				}
			}

			if loop {
				sum++
			}
		}
	}

	return sum
}

// getNextPosition returns the next position based on the current position and direction
func getNextPosition(y, x int, direction Direction) [2]int {
	switch direction {
	case UP:
		return [2]int{y - 1, x}
	case DOWN:
		return [2]int{y + 1, x}
	case LEFT:
		return [2]int{y, x - 1}
	case RIGHT:
		return [2]int{y, x + 1}
	}
	return [2]int{y, x}
}

// isOutsideGrid checks if the position is outside the grid
func isOutsideGrid(pos [2]int, input []string) bool {
	return pos[0] < 0 || pos[1] < 0 || pos[0] >= len(input) || pos[1] >= len(input[0])
}

// turnRight turns the guard to the right
func turnRight(direction Direction) Direction {
	switch direction {
	case UP:
		return RIGHT
	case RIGHT:
		return DOWN
	case DOWN:
		return LEFT
	case LEFT:
		return UP
	}
	return UP
}
