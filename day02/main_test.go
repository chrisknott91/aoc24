package main

import (
	"reflect"
	"testing"
)

func TestParseLine(t *testing.T) {
	tests := []struct {
		input    string
		expected []int
	}{
		{"1 2 3", []int{1, 2, 3}},
		{"4 5 6 7", []int{4, 5, 6, 7}},
		{"10 20 30", []int{10, 20, 30}},
	}

	for _, test := range tests {
		result := parseLine(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("parseLine(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsSafe(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
	}{
		{[]int{7, 6, 4, 2, 1}, true},  // Safe because the levels are all decreasing by 1 or 2.
		{[]int{1, 2, 7, 8, 9}, false}, // Unsafe because 2 7 is an increase of 5.
		{[]int{9, 7, 6, 2, 1}, false}, // Unsafe because 6 2 is a decrease of 4.
		{[]int{1, 3, 2, 4, 5}, false}, // Unsafe because 1 3 is increasing but 3 2 is decreasing.
		{[]int{8, 6, 4, 4, 1}, false}, // Unsafe because 4 4 is neither an increase or a decrease.
		{[]int{1, 3, 6, 7, 9}, true},  // Safe because the levels are all increasing by 1, 2, or 3.
	}

	for _, test := range tests {
		result := isSafe(test.input)
		if result != test.expected {
			t.Errorf("isSafe(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}
