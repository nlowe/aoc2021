package day20

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/nlowe/aoc2021/challenge"
)

func TestA(t *testing.T) {
	t.Skipf("Not solved yet")

	input := challenge.FromLiteral("foobar")

	result := partA(input)

	require.Equal(t, 42, result)
}
