package day10

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/nlowe/aoc2021/challenge"
)

const example = `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`

func TestA(t *testing.T) {
	input := challenge.FromLiteral(example)

	result := partA(input)

	require.Equal(t, 26397, result)
}
