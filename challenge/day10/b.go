package day10

import (
	"fmt"
	"sort"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2021/challenge"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 10, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) int {
	remainders, _ := incomplete(challenge)

	scores := make([]int, len(remainders))
	for i, r := range remainders {
		scores[i] = scoreCompletion(r)
	}

	sort.Ints(scores)
	return scores[len(scores)/2]
}

func scoreCompletion(completion string) (score int) {
	for _, c := range completion {
		score *= 5

		switch c {
		case ')':
			score += 1
		case ']':
			score += 2
		case '}':
			score += 3
		case '>':
			score += 4
		}
	}

	return
}
