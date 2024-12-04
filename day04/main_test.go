package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"AMXSXMAAMM",
		"MSAMASMSMX",
		"XMASAMXAMM",
		"XXAMMXXAMA",
		"SMSMSASXSS",
		"SAXAMASAAA",
		"MAMMMXMMMM",
		"MXMXAXMASX",
	}
	expected := 18

	result := findXMAS(input)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	input := []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"AMXSXMAAMM",
		"MSAMASMSMX",
		"XMASAMXAMM",
		"XXAMMXXAMA",
		"SMSMSASXSS",
		"SAXAMASAAA",
		"MAMMMXMMMM",
		"MXMXAXMASX",
	}
	expected := 9

	result := findMAS(input)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
