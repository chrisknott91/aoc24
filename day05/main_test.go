package main

import (
	"reflect"
	"testing"
)

func TestParseUpdate(t *testing.T) {
	line := "75,47,61,53,29"
	expected := []int{75, 47, 61, 53, 29}
	result := parseUpdate(line)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestParseRule(t *testing.T) {
	line := "47|53"
	expected := [2]int{47, 53}
	result := parseRule(line)

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestParseInput(t *testing.T) {
	input := []string{
		"47|53",
		"97|13",
		"",
		"75,47,61,53,29",
		"97,61,53,29,13",
	}

	expectedRules := map[int][]int{
		47: {53},
		97: {13},
	}
	expectedUpdates := [][]int{
		{75, 47, 61, 53, 29},
		{97, 61, 53, 29, 13},
	}

	rules, updates := parseInput(input)

	if !reflect.DeepEqual(rules, expectedRules) {
		t.Errorf("Expected rules %v, but got %v", expectedRules, rules)
	}

	if !reflect.DeepEqual(updates, expectedUpdates) {
		t.Errorf("Expected updates %v, but got %v", expectedUpdates, updates)
	}
}

func TestSumMiddlePagesValidUpdates(t *testing.T) {
	rules := map[int][]int{
		47: {53, 61, 29, 13},
		97: {13, 61, 47, 75, 29, 53},
		75: {29, 53, 47, 61, 13},
		61: {13, 53, 29},
		53: {29, 13},
		29: {13},
	}
	updates := [][]int{
		{75, 47, 61, 53, 29},
		{97, 61, 53, 29, 13},
		{75, 29, 13},
		{75, 97, 47, 61, 53},
		{61, 13, 29},
		{97, 13, 75, 29, 47},
	}

	validUpdates := filterUpdates(rules, updates, true)
	middleSum := sumMiddlePages(validUpdates)

	expectedSum := 143
	if middleSum != expectedSum {
		t.Errorf("Expected sum of middle pages to be %d, but got %d", expectedSum, middleSum)
	}
}

func TestSumMiddlePagesCorrectedUpdates(t *testing.T) {
	rules := map[int][]int{
		47: {53, 61, 29, 13},
		97: {13, 61, 47, 75, 29, 53},
		75: {29, 53, 47, 61, 13},
		61: {13, 53, 29},
		53: {29, 13},
		29: {13},
	}
	updates := [][]int{
		{75, 47, 61, 53, 29},
		{97, 61, 53, 29, 13},
		{75, 29, 13},
		{75, 97, 47, 61, 53},
		{61, 13, 29},
		{97, 13, 75, 29, 47},
	}

	invalidUpdates := filterUpdates(rules, updates, false)
	correctedUpdates := correctUpdates(rules, invalidUpdates)
	middleSum := sumMiddlePages(correctedUpdates)

	expectedSum := 123
	if middleSum != expectedSum {
		t.Errorf("Expected sum of middle pages to be %d, but got %d", expectedSum, middleSum)
	}
}
