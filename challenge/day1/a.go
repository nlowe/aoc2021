package day1

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2021/challenge"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 1, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

func partA(challenge *challenge.Input) int {
	measurements := challenge.IntSlice()

	count := 0
	last := measurements[0]
	for i := 1; i < len(measurements); i++ {
		current := measurements[i]

		if current > last {
			count++
		}

		last = current
	}

	return count
}
