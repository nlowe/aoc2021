package day13

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
		Short: "Day 13, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

type paper struct {
	grid map[int]map[int]struct{}

	mx int
	my int
}

func (p paper) hasHoleAt(x, y int) bool {
	if g, ok := p.grid[x]; ok {
		if _, ok = g[y]; ok {
			return true
		}
	}

	return false
}

func (p *paper) punch(x, y int) {
	if _, ok := p.grid[x]; !ok {
		p.grid[x] = map[int]struct{}{}
	}

	p.grid[x][y] = struct{}{}
}

func (p paper) countHoles() (result int) {
	for x := 0; x <= p.mx; x++ {
		for y := 0; y <= p.my; y++ {
			if p.hasHoleAt(x, y) {
				result++
			}
		}
	}

	return
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
	if f.along == "x" {
		p.foldX(f.v)
	} else {
		p.foldY(f.v)
	}
}

func (p *paper) foldX(v int) {
	result := map[int]map[int]struct{}{}

	for x := 0; x < v; x++ {
		for y := 0; y <= p.my; y++ {
			if g, ok := p.grid[x]; ok {
				if _, ok := g[y]; ok {
					if _, ok := result[x]; !ok {
						result[x] = map[int]struct{}{}
					}

					result[x][y] = struct{}{}
				}
			}
		}
	}

	for x := v + 1; x <= p.mx; x++ {
		for y := 0; y <= p.my; y++ {
			if g, ok := p.grid[x]; ok {
				if _, ok := g[y]; ok {
					dx := x - v

					if _, ok := result[v-dx]; !ok {
						result[v-dx] = map[int]struct{}{}
					}

					result[v-dx][y] = struct{}{}
				}
			}
		}
	}

	p.grid = result
	p.mx = v
}

func (p *paper) foldY(v int) {
	result := map[int]map[int]struct{}{}

	for x := 0; x <= p.mx; x++ {
		for y := 0; y < v; y++ {

			if g, ok := p.grid[x]; ok {
				if _, ok := g[y]; ok {
					if _, ok := result[x]; !ok {
						result[x] = map[int]struct{}{}
					}

					result[x][y] = struct{}{}
				}
			}
		}
	}

	for x := 0; x <= p.mx; x++ {
		for y := v + 1; y <= p.my; y++ {
			if g, ok := p.grid[x]; ok {
				if _, ok := g[y]; ok {
					dy := y - v

					if _, ok := result[x]; !ok {
						result[x] = map[int]struct{}{}
					}

					result[x][v-dy] = struct{}{}
				}
			}
		}
	}

	p.grid = result
	p.my = v
}

type fold struct {
	along string
	v     int
}

func partA(challenge *challenge.Input) int {
	p, folds := parse(challenge)

	p.apply(folds[0])

	return p.countHoles()
}

func parse(challenge *challenge.Input) (paper, []fold) {
	p := paper{grid: map[int]map[int]struct{}{}}
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

			p.punch(x, y)
		} else {
			parts := strings.Split(strings.TrimPrefix(line, "fold along "), "=")
			v := util.MustAtoI(parts[1])

			folds = append(folds, fold{along: parts[0], v: v})
		}
	}

	return p, folds
}
