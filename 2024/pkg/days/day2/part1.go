package day2

import (
	"aoc_go_2024/pkg/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func SolvePart1() {
	fmt.Println("Solution part 1")

	inputContent := utils.MustHaveFile(utils.GetInputFilePath())
	safeCount := 0

	for _, reportLine := range inputContent {
		level := parseLevels(reportLine)
		var prevSign *int
		broke := false

		for i := 1; i < len(level); i++ {
			diff, sign := absWithSign(level[i] - level[i - 1])
			if diff < 1 || diff > 3 {
				broke = true
				break
			}
			if prevSign != nil && *prevSign != sign {
				broke = true
				break	
			}
			prevSign = &sign
		}

		if (!broke) {
			safeCount++
		}
	}
	fmt.Printf("Safe count is: %d\n", safeCount)
}

func absWithSign(x int) (int, int) {
	if x < 0 {
		return -x, -1
	}
	return x, 1
}

// TODO: into utils? parametrize "parsing" func, here it is into int
func parseLevels(reportLine string) []int {
	reportLevelsSeq := utils.Map(slices.Values(strings.Fields(reportLine)), func(c string) int { return utils.Must(strconv.Atoi(c)) })
	return slices.Collect(reportLevelsSeq)
}