package day13

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2021/challenge"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 13, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: \n%s\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) string {
	p, folds := parse(challenge)

	for _, f := range folds {
		p.apply(f)
	}

	return p.print('\u2588')
}
