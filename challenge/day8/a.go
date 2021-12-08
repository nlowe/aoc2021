package day8

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2021/challenge"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 8, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

func partA(challenge *challenge.Input) int {
	count := 0

	for line := range challenge.Lines() {
		parts := strings.Split(line, " | ")
		outputs := strings.Fields(parts[1])

		// how many segments produce each digit
		one := 2
		four := 4
		seven := 3
		eight := 7

		for _, display := range outputs {
			switch len(display) {
			case one:
				fallthrough
			case four:
				fallthrough
			case seven:
				fallthrough
			case eight:
				count++
			default:
			}
		}
	}

	return count
}
