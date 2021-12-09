package day9

import (
	"fmt"

	"github.com/nlowe/aoc2021/util"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2021/challenge"
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
	risk, _ := findLowPoints(challenge.TileMap())
	return risk
}

func findLowPoints(cave *challenge.TileMap) (risk int, lowPoints []point) {
	w, h := cave.Size()

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			up, uok := cave.TileAt(x, y-1)
			right, rok := cave.TileAt(x+1, y)
			down, dok := cave.TileAt(x, y+1)
			left, lok := cave.TileAt(x-1, y)

			cur, _ := cave.TileAt(x, y)
			c := util.MustAtoI(string(cur))

			if uok && c >= util.MustAtoI(string(up)) {
				continue
			}

			if rok && c >= util.MustAtoI(string(right)) {
				continue
			}

			if dok && c >= util.MustAtoI(string(down)) {
				continue
			}

			if lok && c >= util.MustAtoI(string(left)) {
				continue
			}

			lowPoints = append(lowPoints, point{x, y})
			risk += 1 + c
		}
	}

	return
}
