package day11

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2021/challenge"
	"github.com/nlowe/aoc2021/util"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 11, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

func partA(challenge *challenge.Input) int {
	totalFlashes := 0
	simulate(challenge, func(step, flashes int) bool {
		totalFlashes += flashes

		return step == 99
	})

	return totalFlashes
}

func simulate(challenge *challenge.Input, stop func(step, flashes int) bool) {
	cave := make([]int, 10*10)

	for y, line := range challenge.LineSlice() {
		for x, o := range line {
			cave[y*10+x] = util.MustAtoI(string(o))
		}
	}

	for i := 0; ; i++ {
		for x := 0; x < 10; x++ {
			for y := 0; y < 10; y++ {
				cave[y*10+x]++
			}
		}

		flashes := 0
		for x := 0; x < 10; x++ {
			for y := 0; y < 10; y++ {
				if cave[y*10+x] > 9 {
					flashes += flash(cave, x, y)
				}
			}
		}

		for x := 0; x < 10; x++ {
			for y := 0; y < 10; y++ {
				if cave[y*10+x] == -1 {
					cave[y*10+x] = 0
				}
			}
		}

		if stop(i, flashes) {
			return
		}
	}
}

var delta = []struct {
	x int
	y int
}{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
	{1, 1},
	{1, -1},
	{-1, 1},
	{-1, -1},
}

func flash(cave []int, x, y int) int {
	// Already Flashed
	if cave[y*10+x] == -1 {
		return 0
	}

	flashes := 1
	cave[y*10+x] = -1

	for _, d := range delta {
		// Out of Range
		if x+d.x < 0 || x+d.x >= 10 || y+d.y < 0 || y+d.y >= 10 {
			continue
		}

		// Already Flashed
		if cave[(y+d.y)*10+(x+d.x)] == -1 {
			continue
		}

		cave[(y+d.y)*10+(x+d.x)]++
		if cave[(y+d.y)*10+(x+d.x)] > 9 {
			flashes += flash(cave, x+d.x, y+d.y)
		}
	}

	return flashes
}
