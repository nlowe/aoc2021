package day18

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/nlowe/aoc2021/challenge"
)

func TestB(t *testing.T) {
	t.Skipf("Not solved yet")

	input := challenge.FromLiteral("foobar")

	result := partB(input)

	require.Equal(t, 42, result)
}
