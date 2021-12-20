package day10

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2021/challenge"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 10, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

const openChars = `([{<`

func partA(challenge *challenge.Input) int {
	_, score := incomplete(challenge)

	return score
}

func incomplete(challenge *challenge.Input) (incomplete []string, errorScore int) {
lines:
	for line := range challenge.Lines() {
		var q []rune

		for _, c := range line {
			var expected rune

			//nolint:gocritic // The order of these is actually correct, we're checking if the single rune is present.
			if strings.Contains(openChars, string(c)) {
				q = append([]rune{mirror(c)}, q...)
				continue
			}

			if len(q) == 0 {
				errorScore += errorValue(c)
				continue lines
			}

			expected, q = q[0], q[1:]
			if expected != c {
				errorScore += errorValue(c)
				continue lines
			}
		}

		if len(q) > 0 {
			incomplete = append(incomplete, string(q))
		}
	}

	return
}

func mirror(c rune) rune {
	switch c {
	case '(':
		return ')'
	case '[':
		return ']'
	case '{':
		return '}'
	case '<':
		return '>'
	default:
		panic(fmt.Errorf("no mirror for %s", string(c)))
	}
}

func errorValue(c rune) int {
	switch c {
	case ')':
		return 3
	case ']':
		return 57
	case '}':
		return 1197
	case '>':
		return 25137
	default:
		panic(fmt.Errorf("not a closing character: %s", string(c)))
	}
}
