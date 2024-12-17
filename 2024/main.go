package main

import (
	"aoc_go_2024/pkg/days/day1"
	"aoc_go_2024/pkg/days/day2"
	"aoc_go_2024/pkg/utils"
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
}

func main ()  {
	utils.PrintMsg("Initial")

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
