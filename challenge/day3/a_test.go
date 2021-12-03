package day3

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/nlowe/aoc2021/challenge"
)

const example = `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

func TestA(t *testing.T) {
	input := challenge.FromLiteral(example)

	result := partA(input)

	require.Equal(t, 198, result)
}
