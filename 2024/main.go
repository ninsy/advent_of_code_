package main

import (
	"aoc_go_2024/pkg/days/day1"
	"aoc_go_2024/pkg/days/day2"
	"fmt"
	"os"
)

// NOTE: add new solutions here
var solutions = map[int]Day{
	1: {
		Part1: day1.SolvePart1,
		Part2: day1.SolvePart2,
	},
	2: {
		Part1: day2.SolvePart1,
		Part2: day2.SolvePart2,
	},
	3: {
		Part1: day2.SolvePart1,
		Part2: day2.SolvePart2,
	},
	4: {
		Part1: day2.SolvePart1,
		Part2: day2.SolvePart2,
	},
	5: {
		Part1: day2.SolvePart1,
		Part2: day2.SolvePart2,
	},
	6: {
		Part1: day2.SolvePart1,
		Part2: day2.SolvePart2,
	},
	7: {
		Part1: day2.SolvePart1,
		Part2: day2.SolvePart2,
	},
	8: {
		Part1: day2.SolvePart1,
		Part2: day2.SolvePart2,
	},
	9: {
		Part1: day2.SolvePart1,
		Part2: day2.SolvePart2,
	},
	10: {
		Part1: day2.SolvePart1,
		Part2: day2.SolvePart2,
	},
	11: {
		Part1: day2.SolvePart1,
		Part2: day2.SolvePart2,
	},
	12: {
		Part1: day2.SolvePart1,
		Part2: day2.SolvePart2,
	},
	13: {
		Part1: day2.SolvePart1,
		Part2: day2.SolvePart2,
	},
	14: {
		Part1: day2.SolvePart1,
		Part2: day2.SolvePart2,
	},
	15: {
		Part1: day2.SolvePart1,
		Part2: day2.SolvePart2,
	},
	16: {
		Part1: day2.SolvePart1,
		Part2: day2.SolvePart2,
	},
	17: {
		Part1: day2.SolvePart1,
		Part2: day2.SolvePart2,
	},
	18: {
		Part1: day2.SolvePart1,
		Part2: day2.SolvePart2,
	},
	19: {
		Part1: day2.SolvePart1,
		Part2: day2.SolvePart2,
	},
	20: {
		Part1: day2.SolvePart1,
		Part2: day2.SolvePart2,
	},
	21: {
		Part1: day2.SolvePart1,
		Part2: day2.SolvePart2,
	},
	22: {
		Part1: day2.SolvePart1,
		Part2: day2.SolvePart2,
	},
	23: {
		Part1: day2.SolvePart1,
		Part2: day2.SolvePart2,
	},
	24: {
		Part1: day2.SolvePart1,
		Part2: day2.SolvePart2,
	},
	25: {
		Part1: day2.SolvePart1,
		Part2: day2.SolvePart2,
	},
}

func main ()  {
	argLen := len(os.Args)
	if argLen < 2 || argLen > 3  {
		fmt.Println("Usage: go run main.go <day> [part]")
		return;
	}

	dayStruct := GetDay(os.Args[1])

	switch (argLen) {
	case 2:
		fmt.Println("Running part 1");
		dayStruct.Part1()
		fmt.Println("Running part 2");
		dayStruct.Part2()
	case 3:
		dayStruct.getPart(os.Args[2])();
	}
}
