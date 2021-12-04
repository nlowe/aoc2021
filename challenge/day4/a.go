package day4

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
		Short: "Day 4, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

type slot struct {
	x int
	y int
}

type board struct {
	numbers map[string]slot
	markers map[int]map[int]bool
}

func emptyBoard() board {
	return board{
		numbers: map[string]slot{},
		markers: map[int]map[int]bool{
			0: {},
			1: {},
			2: {},
			3: {},
			4: {},
		},
	}
}

func partA(challenge *challenge.Input) int {
	drawing, boards := parseBoards(challenge)

	for idx, number := range drawing {
		for _, b := range boards {
			if xy, ok := b.numbers[number]; ok {
				b.markers[xy.x][xy.y] = true

				if idx >= 4 && b.won() {
					return b.score(util.MustAtoI(number))
				}
			}
		}
	}

	panic("no solution")
}

func parseBoards(challenge *challenge.Input) ([]string, []board) {
	parts := challenge.LineSlice()
	var boards []board

	drawing := strings.Split(parts[0], ",")

	y := 0
	currentBoard := emptyBoard()
	for i := 2; i < len(parts); i++ {
		if parts[i] == "" {
			boards = append(boards, currentBoard)
			currentBoard = emptyBoard()
			y = 0

			continue
		}

		for x, number := range strings.Fields(parts[i]) {
			currentBoard.numbers[number] = slot{x, y}
		}

		y++
	}

	return drawing, append(boards, currentBoard)
}

func (b board) won() bool {
	// columns
column:
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			if !b.markers[x][y] {
				continue column
			}
		}

		return true
	}

	// rows
row:
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if !b.markers[x][y] {
				continue row
			}
		}

		return true
	}

	return false
}

func (b board) score(called int) int {
	unmarked := 0

	for number, xy := range b.numbers {
		if !b.markers[xy.x][xy.y] {
			unmarked += util.MustAtoI(number)
		}
	}

	return unmarked * called
}
