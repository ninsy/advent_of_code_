package main

import (
	"aoc_go_2024/pkg/days/day1"
	"aoc_go_2024/pkg/days/day10"
	"aoc_go_2024/pkg/days/day11"
	"aoc_go_2024/pkg/days/day12"
	"aoc_go_2024/pkg/days/day13"
	"aoc_go_2024/pkg/days/day14"
	"aoc_go_2024/pkg/days/day15"
	"aoc_go_2024/pkg/days/day16"
	"aoc_go_2024/pkg/days/day17"
	"aoc_go_2024/pkg/days/day18"
	"aoc_go_2024/pkg/days/day19"
	"aoc_go_2024/pkg/days/day2"
	"aoc_go_2024/pkg/days/day20"
	"aoc_go_2024/pkg/days/day21"
	"aoc_go_2024/pkg/days/day22"
	"aoc_go_2024/pkg/days/day23"
	"aoc_go_2024/pkg/days/day24"
	"aoc_go_2024/pkg/days/day25"
	"aoc_go_2024/pkg/days/day3"
	"aoc_go_2024/pkg/days/day4"
	"aoc_go_2024/pkg/days/day5"
	"aoc_go_2024/pkg/days/day6"
	"aoc_go_2024/pkg/days/day7"
	"aoc_go_2024/pkg/days/day8"
	"aoc_go_2024/pkg/days/day9"
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
		Part1: day3.SolvePart1,
		Part2: day3.SolvePart2,
	},
	4: {
		Part1: day4.SolvePart1,
		Part2: day4.SolvePart2,
	},
	5: {
		Part1: day5.SolvePart1,
		Part2: day5.SolvePart2,
	},
	6: {
		Part1: day6.SolvePart1,
		Part2: day6.SolvePart2,
	},
	7: {
		Part1: day7.SolvePart1,
		Part2: day7.SolvePart2,
	},
	8: {
		Part1: day8.SolvePart1,
		Part2: day8.SolvePart2,
	},
	9: {
		Part1: day9.SolvePart1,
		Part2: day9.SolvePart2,
	},
	10: {
		Part1: day10.SolvePart1,
		Part2: day10.SolvePart2,
	},
	11: {
		Part1: day11.SolvePart1,
		Part2: day11.SolvePart2,
	},
	12: {
		Part1: day12.SolvePart1,
		Part2: day12.SolvePart2,
	},
	13: {
		Part1: day13.SolvePart1,
		Part2: day13.SolvePart2,
	},
	14: {
		Part1: day14.SolvePart1,
		Part2: day14.SolvePart2,
	},
	15: {
		Part1: day15.SolvePart1,
		Part2: day15.SolvePart2,
	},
	16: {
		Part1: day16.SolvePart1,
		Part2: day16.SolvePart2,
	},
	17: {
		Part1: day17.SolvePart1,
		Part2: day17.SolvePart2,
	},
	18: {
		Part1: day18.SolvePart1,
		Part2: day18.SolvePart2,
	},
	19: {
		Part1: day19.SolvePart1,
		Part2: day19.SolvePart2,
	},
	20: {
		Part1: day20.SolvePart1,
		Part2: day20.SolvePart2,
	},
	21: {
		Part1: day21.SolvePart1,
		Part2: day21.SolvePart2,
	},
	22: {
		Part1: day22.SolvePart1,
		Part2: day22.SolvePart2,
	},
	23: {
		Part1: day23.SolvePart1,
		Part2: day23.SolvePart2,
	},
	24: {
		Part1: day24.SolvePart1,
		Part2: day24.SolvePart2,
	},
	25: {
		Part1: day25.SolvePart1,
		Part2: day25.SolvePart2,
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
