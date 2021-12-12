package day12

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2021/challenge"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 12, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) int {
	nodes := map[string][]string{}

	for line := range challenge.Lines() {
		parts := strings.Split(line, "-")

		nodes[parts[0]] = append(nodes[parts[0]], parts[1])
		nodes[parts[1]] = append(nodes[parts[1]], parts[0])
	}

	solutions := map[string]struct{}{}

	for n := range nodes {
		if n == "start" || n == "end" || strings.ToLower(n) != n {
			continue
		}

		for _, edge := range nodes["start"] {
			distinctPathsWithVisitTwice(
				nodes,
				edge,
				n,
				false,
				"start",
				solutions,
			)
		}
	}

	return len(solutions)
}

func distinctPathsWithVisitTwice(nodes map[string][]string, start string, canVisitTwice string, alreadyVisitedTwice bool, path string, solutions map[string]struct{}) {
	if start == "end" {
		solutions[path] = struct{}{}
		return
	}

	for _, to := range nodes[start] {
		// Small node?
		secondVisit := alreadyVisitedTwice

		visited := strings.Contains(path, to)
		if visited && strings.ToLower(to) == to {
			if to == canVisitTwice {
				if alreadyVisitedTwice {
					continue
				}

				secondVisit = true
			} else {
				// We've already traversed this node
				continue
			}
		}

		distinctPathsWithVisitTwice(
			nodes,
			to,
			canVisitTwice,
			secondVisit,
			fmt.Sprintf("%s,%s", path, start),
			solutions,
		)
	}
}
