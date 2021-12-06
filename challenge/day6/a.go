package day6

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
		Short: "Day 6, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

func partA(challenge *challenge.Input) int {
	return simulate(challenge, 80)
}

func simulate(challenge *challenge.Input, days int) int {
	raw := strings.Split(challenge.LineSlice()[0], ",")

	fish := make([]int, 9)
	for _, f := range raw {
		fish[util.MustAtoI(f)]++
	}

	for i := 0; i < days; i++ {
		// The head of the counter queue produces new fish today
		bornToday := fish[0]

		// Shuffle everyone down one day, moving today's to the end
		fish = append(fish[1:], fish[0])

		// It takes 6 days for today's fish to recharge
		fish[6] += bornToday

		// and 8 days for the new fish to start producing more fish
		fish[8] = bornToday
	}

	total := 0
	for _, n := range fish {
		total += n
	}

	return total
}
