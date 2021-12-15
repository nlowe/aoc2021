package day15

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/nlowe/aoc2021/challenge"
)

const example = `1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581`

func TestA(t *testing.T) {
	input := challenge.FromLiteral(example)

	result := partA(input)

	require.Equal(t, 40, result)
}
