package day15

import (
	"fmt"
	"strconv"

	"github.com/beefsack/go-astar"
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
	cave := extend(challenge.TileMap())
	w, h := cave.Size()

	_, distance, found := astar.Path(spot{cave: cave, x: 0, y: 0}, spot{cave: cave, x: w - 1, y: h - 1})
	if !found {
		panic("no solution")
	}

	return int(distance)
}

func extend(cave *challenge.TileMap) *challenge.TileMap {
	w, h := cave.Size()

	result := challenge.NewTileMap(w*5, h*5)

	for rptx := 0; rptx < 5; rptx++ {
		for rpty := 0; rpty < 5; rpty++ {
			for x := 0; x < w; x++ {
				for y := 0; y < h; y++ {
					r, _ := cave.TileAt(x, y)

					v := util.MustAtoI(string(r))

					var back rune
					var ok bool
					if rptx > 0 {
						back, ok = result.TileAt(w*(rptx-1)+x, w*rpty+y)
					} else if rpty > 0 {
						back, ok = result.TileAt(w*rptx+x, w*(rpty-1)+y)
					}

					if ok {
						v = util.MustAtoI(string(back)) + 1

						if v == 10 {
							v = 1
						}
					}

					n := rune(strconv.Itoa(v)[0])
					result.SetTile(w*rptx+x, h*rpty+y, n)
				}
			}
		}
	}

	return result
}
