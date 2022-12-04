package day9

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
		Short: "Day 9, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

func partA(challenge *challenge.Input) int {
	risk, _ := findLowPoints(tilemap.FromInputOf[int](challenge, util.MustSingleDigitAToI))
	return risk
}

func findLowPoints(cave *tilemap.TileMap[int]) (risk int, lowPoints []point) {
	w, h := cave.Size()

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			up, uok := cave.TileAt(x, y-1)
			right, rok := cave.TileAt(x+1, y)
			down, dok := cave.TileAt(x, y+1)
			left, lok := cave.TileAt(x-1, y)

			c, _ := cave.TileAt(x, y)

			if uok && c >= up {
				continue
			}

			if rok && c >= right {
				continue
			}

			if dok && c >= down {
				continue
			}

			if lok && c >= left {
				continue
			}

			lowPoints = append(lowPoints, point{x, y})
			risk += 1 + c
		}
	}

	return
}
