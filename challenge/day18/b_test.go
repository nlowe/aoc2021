package day18

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/nlowe/aoc2021/challenge"
)

func TestB(t *testing.T) {
	input := challenge.FromLiteral(example)

	result := partB(input)

	require.Equal(t, 3993, result)
}
