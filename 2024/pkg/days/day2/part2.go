package day2

import (
	"aoc_go_2024/pkg/utils"
	"fmt"
)

func SolvePart2() {
	fmt.Println("Solution part 1")

	inputContent := utils.MustHaveFile(utils.GetInputFilePath())
	safeCount := 0
	processCount := 0

	for _, reportLine := range inputContent {
		level := parseLevels(reportLine)
		wholeOk, _ := testLevel(level)
		if wholeOk {
			safeCount++
			continue
		} else {
			ok := testLevelCombinations(level)
			if ok {
				safeCount++
			}
		}
		processCount++
	}
	fmt.Printf("Safe count is: %d\n", safeCount)
	fmt.Printf("Processed: %d\n", processCount)
}

func splice(s []int, idx int) []int {
	dst := make([]int, len(s))
	copy(dst, s) 
	return append(dst[:idx], dst[idx+1:]...)
}

func testLevelCombinations(level []int) bool {
	for i := 0; i < len(level); i++ {
		currCombination := splice(level, i)
		ok, _ := testLevel(currCombination)
		if ok {
			return true
		}
	}
	return false
}
