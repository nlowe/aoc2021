package day1

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/nlowe/aoc2021/challenge"
)

const testInput = `199
200
208
210
200
207
240
269
260
263`

func TestA(t *testing.T) {
	input := challenge.FromLiteral(testInput)

	result := partA(input)

	require.Equal(t, 7, result)
}
