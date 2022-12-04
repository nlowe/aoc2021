package day15

import (
	"fmt"

	"github.com/nlowe/aoc2021/util/tilemap"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2021/challenge"
	"github.com/nlowe/aoc2021/util"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 15, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

func partA(challenge *challenge.Input) int {
	cave := tilemap.FromInputOf[int](challenge, util.MustSingleDigitAToI)
	w, h := cave.Size()

	cave.CostFunc = tileCost

	_, distance, found := cave.PathBetween(0, 0, w-1, h-1)
	if !found {
		panic("no solution")
	}

	return int(distance)
}

func tileCost(_, to tilemap.TileContainer[int]) float64 {
	return float64(to.Value)
}
