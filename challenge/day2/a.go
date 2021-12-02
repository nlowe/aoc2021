package day2

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2021/challenge"
	"github.com/nlowe/aoc2021/util"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 2, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

func partA(challenge *challenge.Input) int {
	x := 0
	depth := 0

	for instruction := range challenge.Lines() {
		parts := strings.Split(instruction, " ")
		magnitude := util.MustAtoI(parts[1])

		switch parts[0] {
		case "forward":
			x += magnitude
		case "down":
			depth += magnitude
		case "up":
			depth -= magnitude
		}
	}

	return x * depth
}
