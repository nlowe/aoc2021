package day3

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2021/challenge"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 3, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

func partA(challenge *challenge.Input) int {
	lines := challenge.LineSlice()
	gamma := 0
	epsilon := 0

	for idx := range lines[0] {
		most, least := mostLeast(lines, idx)

		gamma <<= 1
		epsilon <<= 1

		if most == '1' {
			gamma |= 1
		}

		if least == '1' {
			epsilon |= 1
		}
	}

	return gamma * epsilon
}

func mostLeast(input []string, idx int) (rune, rune) {
	zeroes := 0
	ones := 0

	for _, line := range input {
		if line[idx] == '0' {
			zeroes++
		} else {
			ones++
		}
	}

	if zeroes > ones {
		return '0', '1'
	}

	return '1', '0'
}
