package day2

import (
	"slices"
	"testing"
)

// TODO: migrate to utils?
func TestPart1(t *testing.T) {
	got := parseLevels("31 28 25 24 23")
	expected := []int{31, 28, 25, 24, 23}
	if !slices.Equal(got, expected) {
		t.Errorf("Slices arent equal: expected: %v, got: %v", expected, got)
	}
}
