package day10

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/nlowe/aoc2021/challenge"
)

func TestB(t *testing.T) {
	input := challenge.FromLiteral(example)

	result := partB(input)

	require.Equal(t, 288957, result)
}

func TestScoreCompletion(t *testing.T) {
	for _, tt := range []struct {
		expected   int
		completion string
	}{
		{expected: 288957, completion: `}}]])})]`},
		{expected: 5566, completion: `)}>]})`},
		{expected: 1480781, completion: `}}>}>))))`},
		{expected: 995444, completion: `]]}}]}]}>`},
		{expected: 294, completion: `])}>`},
	} {
		t.Run(tt.completion, func(t *testing.T) {
			require.Equal(t, tt.expected, scoreCompletion(tt.completion))
		})
	}
}
