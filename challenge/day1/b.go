package day1

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2021/challenge"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 1, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) int {
	measurements := challenge.IntSlice()
	windows := make([]int, 0, len(measurements)/3)

	for i := 0; i < len(measurements)-2; i++ {
		windows = append(windows, measurements[i]+measurements[i+1]+measurements[i+2])
	}

	last := windows[0]
	count := 0

	for i := 1; i < len(windows); i++ {
		if windows[i] > last {
			count++
		}

		last = windows[i]
	}

	return count
}
