package day2

import (
	"slices"
	"testing"
)

// TODO: migrate to utils?
func TestPart2(t *testing.T) {
	testCases := []struct {
		got      []int
		idx 	 int
		expected []int
	}{
		{
			got: []int{31, 28, 25, 24, 23},
			expected: []int{31, 28, 25, 24},
			idx: 4,
		},
		{
			got: []int{31, 28, 25, 24, 23},
			expected: []int{31, 25, 24, 23},
			idx: 1,
		},
	}

	for _, testCase := range testCases {
		r := splice(testCase.got, testCase.idx)

		if !slices.Equal(r, testCase.expected) {
			t.Errorf("Slices arent equal: expected: %v, got: %v", testCase.expected, r)
		}
	}
}
