package day1

import (
	"aoc_go_2024/pkg/utils"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func SolvePart1() {
	fmt.Println("Solution part 1")
	
	numbers := utils.MustHaveFile(utils.GetInputFilePath())

	leftColumnNumbers := make([]int, 0, len(numbers))
	rightColumnNumbers := make([]int, 0, len(numbers))

	for _, pair := range numbers {
		splitPair := strings.Fields(pair)
		left, right := splitPair[0], splitPair[1]
		leftColumnNumbers = append(leftColumnNumbers, utils.Must(strconv.Atoi(left)))
		rightColumnNumbers = append(rightColumnNumbers, utils.Must(strconv.Atoi(right)))
	}

	sort.Slice(leftColumnNumbers, func(i, j int) bool {
		return leftColumnNumbers[i] < leftColumnNumbers[j]
	})

	sort.Slice(rightColumnNumbers, func(i, j int) bool {
		return rightColumnNumbers[i] < rightColumnNumbers[j]
	})

	distance := 0
	for i := 0; i < len(numbers); i++ {
		distance += int(math.Abs(float64(leftColumnNumbers[i] - rightColumnNumbers[i])))
	}

	fmt.Println(distance)
}
