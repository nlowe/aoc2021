package day15

import (
	"fmt"

	"github.com/beefsack/go-astar"
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

type spot struct {
	cave *challenge.TileMap

	x int
	y int
}

func (s spot) PathNeighbors() (results []astar.Pather) {
	for _, delta := range []struct {
		x int
		y int
	}{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	} {
		_, ok := s.cave.TileAt(s.x+delta.x, s.y+delta.y)
		if !ok {
			continue
		}

		results = append(results, spot{cave: s.cave, x: s.x + delta.x, y: s.y + delta.y})
	}

	return
}

func (s spot) PathNeighborCost(to astar.Pather) float64 {
	toSpot := to.(spot)
	r, _ := s.cave.TileAt(toSpot.x, toSpot.y)

	return float64(util.MustAtoI(string(r)))
}

func (s spot) PathEstimatedCost(to astar.Pather) float64 {
	other := to.(spot)
	return float64(util.ManhattanDistance(s.x, s.y, other.x, other.y))
}

var _ astar.Pather = spot{}

func partA(challenge *challenge.Input) int {
	cave := challenge.TileMap()
	w, h := cave.Size()

	_, distance, found := astar.Path(spot{cave: cave, x: 0, y: 0}, spot{cave: cave, x: w - 1, y: h - 1})
	if !found {
		panic("no solution")
	}

	return int(distance)
}
