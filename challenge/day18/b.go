package day18

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2021/challenge"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 18, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) int {
	lines := challenge.LineSlice()

	largest := 0
	for _, a := range lines {
		for _, b := range lines {
			if a == b {
				continue
			}

			aa, _ := Parse(a)
			bb, _ := Parse(b)

			mag := aa.add(bb).magnitude()
			if mag > largest {
				largest = mag
			}

			// Reset pointers and try the other way around
			bb, _ = Parse(b)
			aa, _ = Parse(a)

			mag = bb.add(aa).magnitude()
			if mag > largest {
				largest = mag
			}
		}
	}

	return largest
}
