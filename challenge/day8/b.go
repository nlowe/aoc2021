package day8

import (
	"fmt"
	"sort"
	"strings"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2021/challenge"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 8, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) int {
	count := 0

	for line := range challenge.Lines() {
		parts := strings.Split(line, " | ")
		inputs := strings.Fields(parts[0])
		outputs := strings.Fields(parts[1])

		count += decode(sortAll(inputs), sortAll(outputs))
	}

	return count
}

// decode takes the inputs, determines which lines match which segments, and then returns the base-10
// representation of the outputs, where the 0th element is the most-significant digit.
func decode(inputs []string, outputs []string) int {
	encoding := map[int]string{}

	// For a given 7-segment display with segments labeled as follows:
	//
	//      aaaa
	//     b    c
	//     b    c
	//      dddd
	//     e    f
	//     e    f
	//      gggg
	//
	// The following digits light the following segments:
	//
	//     0: a, b, c,    e, f, g
	//     1:       c,       f
	//     2: a,    c, d, e,    g
	//     3: a,    c, d,    f, g
	//     4:    b, c, d,    f
	//     5: a, b,    d,    f, g
	//     6: a, b,    d, e, f, g
	//     7: a,    c,       f
	//     8: a, b, c, d, e, f, g
	//     9: a, b, c, d,    f, g

	// We can immediately solve for 1, 4, 7, and 8 as they use a unique number of segments:
	for _, option := range inputs {
		switch len(option) {
		case 2:
			encoding[1] = option
		case 3:
			encoding[7] = option
		case 4:
			encoding[4] = option
		case 7:
			encoding[8] = option
		}
	}

	// Next, we can compare the well-known segments to 3, 6, 9, and 0 and solve for them
	// based on what segments are shared:
	for _, segments := range inputs {
		if len(segments) == 5 && len(leftBisection(encoding[1], segments)) == 0 {
			// Solve for 3. We can't differentiate between 2 and 5 just yet
			encoding[3] = segments
		} else if len(segments) == 6 {
			switch {
			// Solve for 0, 6, and 9
			case len(leftBisection(encoding[1], segments)) == 1:
				// 1 has exactly 1 unshared segment with 6
				encoding[6] = segments
			case len(leftBisection(encoding[4], segments)) == 0:
				// 4 shares all of its segments with 9
				encoding[9] = segments
			case len(leftBisection(encoding[4], segments)) == 1:
				// 4 has exactly 1 unshared segment with 0
				encoding[0] = segments
			}
		}
	}

	// Now that we know what 6 is, we can find the actual top-right segment by finding the segment
	// lit by a 1 but not by a 6. We can then use this segment to differentiate between a 2 (which
	// will have this segment lit) and a 5 (which will not):
	tr := leftBisection(encoding[1], encoding[6])[0]
	for _, segments := range inputs {
		if len(segments) != 5 || segments == encoding[3] {
			continue
		}

		if strings.Contains(segments, string(tr)) {
			encoding[2] = segments
		} else {
			encoding[5] = segments
		}
	}

	// Build up the result from the decoded outputs
	result := 0
	for _, display := range outputs {
		result *= 10
		for k, v := range encoding {
			if v == display {
				result += k
			}
		}
	}

	return result
}

// leftBisection returns a string made up of characters from a not present in b
func leftBisection(a, b string) string {
	result := strings.Builder{}

search:
	for _, aa := range a {
		for _, bb := range b {
			if aa == bb {
				continue search
			}
		}

		result.WriteRune(aa)
	}

	return result.String()
}

// sortAll returns s where each element in s sorted in-place
func sortAll(s []string) []string {
	result := make([]string, len(s))

	for i, str := range s {
		v := []rune(str)
		sort.Slice(v, func(i, j int) bool {
			return v[i] < v[j]
		})

		result[i] = string(v)
	}

	return result
}
