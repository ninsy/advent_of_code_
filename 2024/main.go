package main

import (
	"aoc_go_2024/pkg/days/day1"
	"aoc_go_2024/pkg/days/day2"
	"aoc_go_2024/pkg/utils"
	"fmt"
	"os"
	"strconv"
)

type Day struct {
	Part1 func () 
	Part2 func ()
}

func (d *Day) getPart (partStr string) func() {
	part := utils.Must(strconv.Atoi(partStr))
	if (part == 1) {
		fmt.Println("Running part 1");
		return d.Part1
	} else if (part == 2) {
		fmt.Println("Running part 2");
		return d.Part2
	}
	panic(fmt.Errorf("invalid part: either 1 or 2"))
}

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

	dayStruct := getDay(os.Args[1])

	switch (argLen) {
	case 2:
		fmt.Println("Running part 1");
		dayStruct.Part1()
		fmt.Println("Running part 2");
		dayStruct.Part1()
	case 3:
		dayStruct.getPart(os.Args[2])();
	}
}

func getDay(dayStr string) Day {
	d := utils.Must(strconv.Atoi(dayStr))
	return utils.MustHaveKey(solutions, d)
}
