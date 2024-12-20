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
	// TODO: doesnt remove faulty properly?
	for _, reportLine := range inputContent {
		level := parseLevels(reportLine)

		ok, idx := testLevel(level)
		if (ok) {
			// fmt.Println("Is ok: ", level)
			safeCount++
		} else {
			fmt.Println("Before 'splice': ", level)
			withoutFaulty := append(level[:idx], level[idx+1:]...)
			nowOk, _ := testLevel(withoutFaulty)
			if nowOk {
				// fmt.Println("Is ok: ", withoutFaulty)
				safeCount++
			} else {
				fmt.Println("Fucked up anyway:", withoutFaulty)
			}
		}
		processCount++
	}
	fmt.Printf("Safe count is: %d\n", safeCount)
	fmt.Printf("Processed: %d\n", processCount)
}

func splice(s []int, idx int) []int {
	return append(s[:idx], s[idx+1:]...)
}

// TODO: it doesnt cover case when item at idx 0 fucks up everything
// TODO: would need to iterate over whole level, remove item, check if it runs, count failed seq, if failed seq > 1 -> asssume row is failed 

func testLevel(level []int) (bool, int) {
	var prevSign *int = nil
	for i := 1; i < len(level); i++ {
		diff, sign := absWithSign(level[i] - level[i - 1])
		if diff < 1 || diff > 3 {
			return false, i
		}
		if prevSign != nil && *prevSign != sign {
			return false, i
		}
		prevSign = &sign
	}
	return true, -1
}
