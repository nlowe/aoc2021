package day7

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
		Short: "Day 7, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

func partA(challenge *challenge.Input) int {
	return minimize(challenge, func(n int) int {
		return n
	})
}

func minimize(challenge *challenge.Input, score func(int) int) int {
	rawPositions := strings.Split(<-challenge.Lines(), ",")

	var min, max int

	crabs := make([]int, len(rawPositions))
	for i, position := range rawPositions {
		crabs[i] = util.MustAtoI(position)

		if i == 0 {
			min = crabs[i]
			max = crabs[i]
		} else {
			if crabs[i] > max {
				max = crabs[i]
			}

			if crabs[i] < min {
				min = crabs[i]
			}
		}
	}

	minCost := 0
	for goal := min; goal <= max; goal++ {
		cost := 0

		for _, crab := range crabs {
			cost += score(util.IntAbs(goal - crab))
		}

		if goal == min || cost < minCost {
			minCost = cost
		}
	}

	return minCost
}
