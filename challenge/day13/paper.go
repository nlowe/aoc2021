package day13

import (
	"strings"

	"github.com/nlowe/aoc2021/challenge"
	"github.com/nlowe/aoc2021/util"
)

type point struct {
	x int
	y int
}

type fold struct {
	along string
	v     int
}

type paper struct {
	holes []point

	mx int
	my int
}

func parse(challenge *challenge.Input) (paper, []fold) {
	p := paper{}
	var folds []fold

	parsingPoints := true
	for line := range challenge.Lines() {
		if line == "" {
			parsingPoints = false

			continue
		}

		if parsingPoints {
			parts := strings.Split(line, ",")

			x := util.MustAtoI(parts[0])
			y := util.MustAtoI(parts[1])
			p.mx = util.IntMax(p.mx, x)
			p.my = util.IntMax(p.my, y)

			p.holes = append(p.holes, point{x, y})
		} else {
			parts := strings.Split(strings.TrimPrefix(line, "fold along "), "=")
			v := util.MustAtoI(parts[1])

			folds = append(folds, fold{along: parts[0], v: v})
		}
	}

	return p, folds
}

func (p paper) hasHoleAt(x, y int) bool {
	for _, hole := range p.holes {
		if hole.x == x && hole.y == y {
			return true
		}
	}

	return false
}

func (p paper) countHoles() (result int) {
	seen := map[point]struct{}{}

	for _, hole := range p.holes {
		seen[hole] = struct{}{}
	}

	return len(seen)
}

func (p paper) print(fill rune) string {
	result := strings.Builder{}

	for y := 0; y <= p.my; y++ {
		for x := 0; x <= p.mx; x++ {
			if p.hasHoleAt(x, y) {
				result.WriteRune(fill)
			} else {
				result.WriteRune(' ')
			}
		}

		result.WriteString("\n")
	}

	return result.String()
}

func (p *paper) apply(f fold) {
	for i, hole := range p.holes {
		switch {
		case f.along == "x" && hole.x > f.v:
			p.holes[i] = point{f.v - (hole.x - f.v), hole.y}
		case f.along == "y" && hole.y > f.v:
			p.holes[i] = point{hole.x, f.v - (hole.y - f.v)}
		}
	}

	if f.along == "x" {
		p.mx = f.v
	} else {
		p.my = f.v
	}
}
