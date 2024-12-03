package main

import (
	"reflect"
	"testing"
)

func TestParseLine(t *testing.T) {
	line := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	expected := [][]int{{2, 4}, {0, 0}, {8, 5}}
	result, _ := parseLine(line, true)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("parseLine(%q) = %v; want %v", line, result, expected)
	}
}
