package main

import (
	"reflect"
	"testing"
)

func TestSplitInput(t *testing.T) {
	input := []string{
		"3 4",
		"4 3",
		"2 5",
		"1 3",
		"3 9",
		"3 3",
	}

	expectedLeft := []int{3, 4, 2, 1, 3, 3}
	expectedRight := []int{4, 3, 5, 3, 9, 3}

	left, right := splitInput(input)

	if !reflect.DeepEqual(left, expectedLeft) {
		t.Errorf("expected left = %v, got %v", expectedLeft, left)
	}

	if !reflect.DeepEqual(right, expectedRight) {
		t.Errorf("expected right = %v, got %v", expectedRight, right)
	}
}
func TestSimilarityScore(t *testing.T) {
	left := []int{3, 4, 2, 1, 3, 3}
	right := []int{4, 3, 5, 3, 9, 3}

	expected := 31
	result := similarityScore(left, right)

	if result != expected {
		t.Errorf("expected similarity score = %d, got %d", expected, result)
	}
}

func TestTotalDistance(t *testing.T) {
	left := []int{3, 4, 2, 1, 3, 3}
	right := []int{4, 3, 5, 3, 9, 3}

	expected := 11
	result := totalDistance(left, right)

	if result != expected {
		t.Errorf("expected total distance = %d, got %d", expected, result)
	}
}
