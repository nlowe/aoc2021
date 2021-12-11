package day11

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2021/challenge"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 11, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) int {
	answer := 0
	simulate(challenge, func(step, flashes int) bool {
		if flashes == 100 {
			answer = step + 1
			return true
		}

		return false
	})

	return answer
}
