package day1

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/nlowe/aoc2021/challenge"
)

func TestB(t *testing.T) {
	input := challenge.FromLiteral(testInput)

	result := partB(input)

	require.Equal(t, 5, result)
}
