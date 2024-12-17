package day1

import (
	"aoc_go_2024/pkg/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func SolvePart2() {
	fmt.Println("Solution part 2")
	numbers := utils.MustHaveFile(utils.GetInputFilePath())

	leftColumnNumbers := make([]int, 0, len(numbers))
	
	rightColumnOccurences := make(map[int]int)
	
	rightColumnOccurences[1] = 1;

	for _, pair := range numbers {		
		splitPairSeq := utils.Map(slices.Values(strings.Fields(pair)), func(c string) int { return utils.Must(strconv.Atoi(c)) })
		splitPair := slices.Collect(splitPairSeq)

		left, right := splitPair[0], splitPair[1]
		leftColumnNumbers = append(leftColumnNumbers, left)
		rightColumnOccurences[right] += 1
	}

	similarity := 0
	for _, num := range leftColumnNumbers {
		similarity += num * rightColumnOccurences[num]
	}

	fmt.Printf("Similiary is: %d\n", similarity)
}