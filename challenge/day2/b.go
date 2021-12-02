package day2

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2021/challenge"
	"github.com/nlowe/aoc2021/util"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 2, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) int {
	aim := 0
	x := 0
	depth := 0

	for instruction := range challenge.Lines() {
		parts := strings.Split(instruction, " ")
		magnitude := util.MustAtoI(parts[1])

		switch parts[0] {
		case "forward":
			x += magnitude
			depth += aim * magnitude
		case "down":
			aim += magnitude
		case "up":
			aim -= magnitude
		}
	}

	return x * depth
}
