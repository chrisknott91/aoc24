package main

import (
	"testing"
)

func TestCalculateVisitedPositions_EmptyGrid(t *testing.T) {
	input := []string{
		"..........",
		"..........",
		"..........",
		"..........",
		"..........",
		"..........",
		"..........",
		"..........",
		"..........",
		"..........",
	}

	expected := 1
	result := calculateVisitedPositions(input)

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestCalculateVisitedPositions_FullObstacles(t *testing.T) {
	input := []string{
		"##########",
		"##########",
		"##########",
		"##########",
		"##########",
		"##########",
		"##########",
		"##########",
		"##########",
		"#####^####",
	}

	expected := 1
	result := calculateVisitedPositions(input)

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestFindLoopCausingObstructions_NoObstacles(t *testing.T) {
	input := []string{
		"..........",
		"..........",
		"..........",
		"..........",
		"..........",
		"..........",
		"..........",
		"..........",
		"..........",
		".....^....",
	}

	expected := 0
	result := findLoopCausingObstructions(input)

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestCalculateVisitedPositions_CustomGrid(t *testing.T) {
	input := []string{
		"....#.....",
		".........#",
		"..........",
		"..#.......",
		".......#..",
		"..........",
		".#..^.....",
		"........#.",
		"#.........",
		"......#...",
	}

	expected := 41
	result := calculateVisitedPositions(input)

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestFindLoopCausingObstructions_CustomGrid(t *testing.T) {
	input := []string{
		"....#.....",
		".........#",
		"..........",
		"..#.......",
		".......#..",
		"..........",
		".#..^.....",
		"........#.",
		"#.........",
		"......#...",
	}

	expected := 6
	result := findLoopCausingObstructions(input)

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
