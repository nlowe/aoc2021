package day5

import (
	"strings"

	"github.com/nlowe/aoc2021/challenge"
	"github.com/nlowe/aoc2021/util"
)

type board map[int]map[int]int

func (b board) Ingest(input *challenge.Input, filter func(x1, x2, y1, y2 int) bool) {
	for l := range input.Lines() {
		parts := strings.Split(l, " -> ")

		start := strings.Split(parts[0], ",")
		end := strings.Split(parts[1], ",")

		x1 := util.MustAtoI(start[0])
		y1 := util.MustAtoI(start[1])

		x2 := util.MustAtoI(end[0])
		y2 := util.MustAtoI(end[1])

		if filter(x1, x2, y1, y2) {
			b.marchAndMark(x1, y1, x2, y2)
		}
	}
}

func (b board) marchAndMark(x1, y1, x2, y2 int) {
	var dx, dy int
	if x1 > x2 {
		dx = -1
	} else if x1 < x2 {
		dx = 1
	}

	if y1 > y2 {
		dy = -1
	} else if y1 < y2 {
		dy = 1
	}

	x := x1
	y := y1

	for {
		if _, ok := b[x]; !ok {
			b[x] = map[int]int{}
		}

		b[x][y]++

		if x == x2 && y == y2 {
			break
		}

		x += dx
		y += dy
	}
}

func (b board) overlappingPoints() (candidates int) {
	for _, column := range b {
		for _, count := range column {
			if count > 1 {
				candidates++
			}
		}
	}

	return candidates
}
