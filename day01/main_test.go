package main

import (
	"os"
	"testing"
)

func TestReadInput(t *testing.T) {
	// Create a temporary file with test data
	file, err := os.CreateTemp("", "test_input.txt")
	if err != nil {
		t.Fatalf("error creating temp file: %v", err)
	}
	defer os.Remove(file.Name())

	// Write test data to the file
	file.WriteString("3 4\n4 3\n2 5\n1 3\n3 9\n3 3\n")
	file.Close()

	left, right := readInput(file.Name())

	expectedLeft := []int{3, 4, 2, 1, 3, 3}
	expectedRight := []int{4, 3, 5, 3, 9, 3}

	for i, v := range left {
		if v != expectedLeft[i] {
			t.Errorf("expected left[%d] = %d, got %d", i, expectedLeft[i], v)
		}
	}

	for i, v := range right {
		if v != expectedRight[i] {
			t.Errorf("expected right[%d] = %d, got %d", i, expectedRight[i], v)
		}
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
