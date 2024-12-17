package main

import (
	"aoc_go_2024/pkg/utils"
	"fmt"
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


func GetDay(dayStr string) Day {
	d := utils.Must(strconv.Atoi(dayStr))
	return utils.MustHaveKey(solutions, d)
}
