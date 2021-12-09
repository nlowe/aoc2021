package day9

import (
	"fmt"
	"sort"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2021/challenge"
	"github.com/nlowe/aoc2021/util"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 9, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

type point struct {
	x int
	y int
}

func partB(challenge *challenge.Input) int {
	cave := challenge.TileMap()

	_, lowPoints := findLowPoints(cave)

	sizes := make([]int, len(lowPoints))
	for i, p := range lowPoints {
		sizes[i] = basinSize(cave, p)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))
	return sizes[0] * sizes[1] * sizes[2]
}

func basinSize(m *challenge.TileMap, low point) int {
	// Keep track of points in this basin, starting with the low point
	seen := map[point]struct{}{low: {}}

	// Check all points surrounding this low point
	toCheck := []point{
		{x: low.x, y: low.y - 1},
		{x: low.x + 1, y: low.y},
		{x: low.x, y: low.y + 1},
		{x: low.x - 1, y: low.y},
	}

	size := 1

	var p point
	for len(toCheck) > 0 {
		p, toCheck = toCheck[0], toCheck[1:]

		// Have we already seen this point?
		if _, alreadyProcessed := seen[p]; alreadyProcessed {
			continue
		}

		// Is this point outside the cave?
		pv, ok := m.TileAt(p.x, p.y)
		if !ok {
			continue
		}

		// Any point with a height 9 does not belong to any basin
		v := util.MustAtoI(string(pv))
		if v == 9 {
			continue
		}

		// Remember this point
		seen[p] = struct{}{}
		size++

		// And check its neighbors...
		toCheck = append(toCheck, []point{
			{x: p.x, y: p.y - 1},
			{x: p.x + 1, y: p.y},
			{x: p.x, y: p.y + 1},
			{x: p.x - 1, y: p.y},
		}...)
	}

	return size
}
