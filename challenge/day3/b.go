package day3

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2021/challenge"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 3, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) int {
	lines := challenge.LineSlice()

	return search(lines, oxygenFilter) * search(lines, co2Filter)
}

func oxygenFilter(most, _ rune) rune {
	return most
}

func co2Filter(_, least rune) rune {
	return least
}

func search(input []string, filter func(rune, rune) rune) int {
	remaining := make([]string, len(input))
	copy(remaining, input)

	offset := 0

	for len(remaining) > 1 {
		left := make([]string, 0, len(remaining))
		target := filter(mostLeast(remaining, offset))

		for _, line := range remaining {
			if rune(line[offset]) == target {
				left = append(left, line)
			}
		}

		remaining = left
		offset++
	}

	v, _ := strconv.ParseInt(remaining[0], 2, 32)
	return int(v)
}
