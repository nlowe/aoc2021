package day9

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/nlowe/aoc2021/challenge"
)

const example = `2199943210
3987894921
9856789892
8767896789
9899965678`

func TestA(t *testing.T) {
	input := challenge.FromLiteral(example)

	result := partA(input)

	require.Equal(t, 15, result)
}
