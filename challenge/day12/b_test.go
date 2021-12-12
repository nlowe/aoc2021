package day12

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/nlowe/aoc2021/challenge"
)

func TestB(t *testing.T) {
	for _, tt := range []struct {
		name     string
		edges    string
		expected int
	}{
		{name: "small", edges: smallExample, expected: 36},
		{name: "medium", edges: mediumExample, expected: 103},
		{name: "large", edges: largeExample, expected: 3509},
	} {
		t.Run(tt.name, func(t *testing.T) {
			input := challenge.FromLiteral(tt.edges)

			result := partB(input)

			require.Equal(t, tt.expected, result)
		})
	}
}
