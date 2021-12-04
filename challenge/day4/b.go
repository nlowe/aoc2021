package day4

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2021/challenge"
	"github.com/nlowe/aoc2021/util"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 4, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) int {
	drawing, boards := parseBoards(challenge)

	for _, number := range drawing {
		remainingBoards := make([]board, 0, len(boards))

		for _, b := range boards {
			if xy, ok := b.numbers[number]; ok {
				b.markers[xy.x][xy.y] = true

				if !b.won() {
					remainingBoards = append(remainingBoards, b)
				} else if len(boards) == 1 {
					return b.score(util.MustAtoI(number))
				}
			} else {
				remainingBoards = append(remainingBoards, b)
			}
		}

		boards = remainingBoards
	}

	panic("no solution")
}
