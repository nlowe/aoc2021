package day14

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2021/challenge"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 14, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

func partA(challenge *challenge.Input) uint64 {
	return react(challenge, 10)
}

func react(challenge *challenge.Input, rounds int) uint64 {
	lines := challenge.LineSlice()

	p := lines[0]
	rules := map[string]string{}
	for i := 2; i < len(lines); i++ {
		parts := strings.Split(lines[i], " -> ")
		rules[parts[0]] = parts[1]
	}

	// Instead of keeping track of the whole polymer chain, just keep track of pairs
	// starting with the initial template.
	pairs := map[string]uint64{}
	for i := 0; i < len(p)-1; i++ {
		pairs[p[i:i+2]]++
	}

	for i := 0; i < rounds; i++ {
		nextChain := map[string]uint64{}

		// Count Pars produced by the replacement rules
		for pair, count := range pairs {
			nextChain[string(pair[0])+rules[pair]] += count
			nextChain[rules[pair]+string(pair[1])] += count
		}

		pairs = nextChain
	}

	// The final count is produced by counting the first element of each pair
	finalPolymer := map[string]uint64{}
	for pair, v := range pairs {
		finalPolymer[string(pair[0])] += v
	}

	// Also count the last element of the template polymer, since it always
	// remains at the end.
	finalPolymer[string(p[len(p)-1])]++

	return ans(finalPolymer)
}

func ans(pairs map[string]uint64) uint64 {
	most := uint64(0)
	least := uint64(2<<63 - 1)
	for _, v := range pairs {
		if v > most {
			most = v
		}

		if v < least {
			least = v
		}
	}

	return most - least
}
