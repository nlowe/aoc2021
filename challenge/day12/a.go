package day12

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2021/challenge"
)

const (
	start = "start"
	end   = "end"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 12, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

func partA(challenge *challenge.Input) int {
	nodes := map[string][]string{}

	for line := range challenge.Lines() {
		parts := strings.Split(line, "-")

		nodes[parts[0]] = append(nodes[parts[0]], parts[1])
		nodes[parts[1]] = append(nodes[parts[1]], parts[0])
	}

	count := 0

	for _, edge := range nodes[start] {
		count += distinctPaths(nodes, edge, map[string]struct{}{start: {}})
	}

	return count
}

func distinctPaths(nodes map[string][]string, startAt string, seen map[string]struct{}) (count int) {
	if startAt == end {
		return 1
	}

	seen[startAt] = struct{}{}
	for _, to := range nodes[startAt] {
		// Small node?
		_, visited := seen[to]
		if visited && strings.ToLower(to) == to {
			// We've already traversed this node
			continue
		}

		clonedSeen := map[string]struct{}{}
		for k, v := range seen {
			clonedSeen[k] = v
		}

		count += distinctPaths(nodes, to, clonedSeen)
	}

	return count
}
