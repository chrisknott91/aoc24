package main

import (
	"fmt"
	"github.com/chrisknott91/aoc24/utils"
)

func main() {
	var input = utils.ReadInput("inputs/day04.txt")
	fmt.Println("Number of accuring XMAS: ", findXMAS(input))
	fmt.Println("Number of accuring X-MAS: ", findMAS(input))
}

func findXMAS(input []string) int {
	sum := 0
	for i := range input {
		for j := range input[i] {
			if input[i][j] == 'X' {
				sum += checkNeighbours(input, i, j)
			}
		}
	}
	return sum
}

func findMAS(input []string) int {
	sum := 0
	for i := range input {
		for j := range input[i] {
			if input[i][j] == 'A' {
				sum += checkNeighbours2(input, i, j)
			}
		}
	}
	return sum
}

func checkNeighbours(input []string, i, j int) int {
	sum := 0
	// left
	if j > 2 && input[i][j-1] == 'M' && input[i][j-2] == 'A' && input[i][j-3] == 'S' {
		sum++
	}
	// right
	if j < len(input[0])-3 && input[i][j+1] == 'M' && input[i][j+2] == 'A' && input[i][j+3] == 'S' {
		sum++
	}
	// up
	if i > 2 && input[i-1][j] == 'M' && input[i-2][j] == 'A' && input[i-3][j] == 'S' {
		sum++
	}
	// down
	if i < len(input)-3 && input[i+1][j] == 'M' && input[i+2][j] == 'A' && input[i+3][j] == 'S' {
		sum++
	}
	// left-up
	if i > 2 && j > 2 && input[i-1][j-1] == 'M' && input[i-2][j-2] == 'A' && input[i-3][j-3] == 'S' {
		sum++
	}
	// right-up
	if i > 2 && j < len(input[0])-3 && input[i-1][j+1] == 'M' && input[i-2][j+2] == 'A' && input[i-3][j+3] == 'S' {
		sum++
	}
	// left-down
	if i < len(input)-3 && j > 2 && input[i+1][j-1] == 'M' && input[i+2][j-2] == 'A' && input[i+3][j-3] == 'S' {
		sum++
	}
	// right-down
	if i < len(input)-3 && j < len(input[0])-3 && input[i+1][j+1] == 'M' && input[i+2][j+2] == 'A' && input[i+3][j+3] == 'S' {
		sum++
	}
	return sum
}

func checkNeighbours2(input []string, i, j int) int {
	sum := 0
	// M Left-up M Right-up
	if i > 0 && j > 0 && j < len(input[0])-1 && i < len(input)-1 &&
		i+1 < len(input) && i-1 >= 0 && j+1 < len(input[0]) && j-1 >= 0 &&
		input[i-1][j-1] == 'M' &&
		input[i+1][j-1] == 'M' &&
		input[i+1][j+1] == 'S' &&
		input[i-1][j+1] == 'S' {
		sum++
	}
	// M Left-up S Right-up
	if i > 0 && j > 0 && j < len(input[0])-1 && i < len(input)-1 &&
		i+1 < len(input) && i-1 >= 0 && j+1 < len(input[0]) && j-1 >= 0 &&
		input[i-1][j-1] == 'M' &&
		input[i+1][j-1] == 'S' &&
		input[i+1][j+1] == 'S' &&
		input[i-1][j+1] == 'M' {
		sum++
	}
	// S Left-up S Right-up
	if i > 0 && j > 0 && j < len(input[0])-1 && i < len(input)-1 &&
		i+1 < len(input) && i-1 >= 0 && j+1 < len(input[0]) && j-1 >= 0 &&
		input[i-1][j-1] == 'S' &&
		input[i+1][j-1] == 'S' &&
		input[i+1][j+1] == 'M' &&
		input[i-1][j+1] == 'M' {
		sum++
	}
	// S Left-up M Right-up
	if i > 0 && j > 0 && j < len(input[0])-1 && i < len(input)-1 &&
		i+1 < len(input) && i-1 >= 0 && j+1 < len(input[0]) && j-1 >= 0 &&
		input[i-1][j-1] == 'S' &&
		input[i+1][j-1] == 'M' &&
		input[i+1][j+1] == 'M' &&
		input[i-1][j+1] == 'S' {
		sum++
	}
	return sum
}
