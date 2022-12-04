package day15

import (
	"fmt"

	"github.com/nlowe/aoc2021/util/tilemap"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2021/challenge"
	"github.com/nlowe/aoc2021/util"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 15, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) int {
	cave := extend(tilemap.FromInputOf[int](challenge, util.MustSingleDigitAToI))
	w, h := cave.Size()

	cave.CostFunc = tileCost

	_, distance, found := cave.PathBetween(0, 0, w-1, h-1)
	if !found {
		panic("no solution")
	}

	return int(distance)
}

func extend(cave *tilemap.TileMap[int]) *tilemap.TileMap[int] {
	w, h := cave.Size()

	result := tilemap.Of[int](w*5, h*5)

	for rptx := 0; rptx < 5; rptx++ {
		for rpty := 0; rpty < 5; rpty++ {
			for x := 0; x < w; x++ {
				for y := 0; y < h; y++ {
					v, _ := cave.TileAt(x, y)

					var back int
					var ok bool
					if rptx > 0 {
						back, ok = result.TileAt(w*(rptx-1)+x, w*rpty+y)
					} else if rpty > 0 {
						back, ok = result.TileAt(w*rptx+x, w*(rpty-1)+y)
					}

					if ok {
						v = back + 1

						if v == 10 {
							v = 1
						}
					}

					result.SetTile(w*rptx+x, h*rpty+y, v%10)
				}
			}
		}
	}

	return result
}
