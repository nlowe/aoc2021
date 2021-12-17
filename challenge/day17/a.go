package day17

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
		Short: "Day 17, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

func partA(challenge *challenge.Input) int {
	var best int
	search(challenge, func(top int) {
		if top > best {
			best = top
		}
	})

	return best
}

func search(challenge *challenge.Input, hit func(top int)) {
	xmin, xmax, ymin, ymax := target(challenge)

	// TODO: 500 seems to be sufficient for my input and the tests.
	//       Can we solve this with IK or otherwise figure out the limit?
	for vx := 0; vx < 500; vx++ {
		for vy := ymin; vy < 500; vy++ {
			top, ok := test(xmin, xmax, ymin, ymax, vx, vy)

			if ok {
				hit(top)
			}
		}
	}
}

func target(challenge *challenge.Input) (int, int, int, int) {
	raw := strings.TrimPrefix(<-challenge.Lines(), "target area: ")
	parts := strings.Split(raw, ", ")

	rawX := strings.Split(strings.TrimPrefix(parts[0], "x="), "..")
	rawY := strings.Split(strings.TrimPrefix(parts[1], "y="), "..")

	x1 := util.MustAtoI(rawX[0])
	x2 := util.MustAtoI(rawX[1])
	xmin := util.IntMin(x1, x2)
	xmax := util.IntMax(x1, x2)

	y1 := util.MustAtoI(rawY[0])
	y2 := util.MustAtoI(rawY[1])
	ymin := util.IntMin(y1, y2)
	ymax := util.IntMax(y1, y2)

	return xmin, xmax, ymin, ymax
}

func test(xmin, xmax, ymin, ymax, vx, vy int) (int, bool) {
	var x, y, maxHeight int

	for {
		x += vx
		y += vy

		if y > maxHeight {
			maxHeight = y
		}

		if vx < 0 {
			vx++
		} else if vx > 0 {
			vx--
		}

		vy--

		if x >= xmin && x <= xmax && y >= ymin && y <= ymax {
			return maxHeight, true
		}

		// already past the bounding box
		if y < ymin || (vx == 0 && (x > xmax || x < xmin)) {
			return maxHeight, false
		}
	}
}
