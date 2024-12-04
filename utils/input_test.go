package utils

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
	file.WriteString("line1\nline2\nline3\n")
	file.Close()

	// Use ReadInput to read the file
	input := ReadInput(file.Name())

	expected := []string{"line1", "line2", "line3"}

	if len(input) != len(expected) {
		t.Fatalf("expected %d lines, got %d", len(expected), len(input))
	}

	for i, line := range input {
		if line != expected[i] {
			t.Errorf("expected line %d to be %q, got %q", i, expected[i], line)
		}
	}
}
