package day18

import (
	"testing"

	"github.com/nlowe/aoc2021/challenge"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	for _, tt := range []struct {
		line string
	}{
		{line: "[1,2]"},
		{line: "[[1,2],3]"},
		{line: "[9,[8,7]]"},
		{line: "[[1,9],[8,5]]"},
		{line: "[[[[1,2],[3,4]],[[5,6],[7,8]]],9]"},
		{line: "[[[9,[3,8]],[[0,9],6]],[[[3,7],[4,9]],3]]"},
		{line: "[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]"},
	} {
		t.Run(tt.line, func(t *testing.T) {
			sut, remaining := Parse(tt.line)

			assert.Empty(t, remaining)
			assert.Equal(t, tt.line, sut.String())
		})
	}
}

func TestAddAll(t *testing.T) {
	for _, tt := range []struct {
		input    string
		expected string
	}{
		{
			input: `[1,1]
[2,2]
[3,3]
[4,4]`,
			expected: "[[[[1,1],[2,2]],[3,3]],[4,4]]",
		},
		{
			input: `[1,1]
[2,2]
[3,3]
[4,4]
[5,5]`,
			expected: "[[[[3,0],[5,3]],[4,4]],[5,5]]",
		},
		{
			input: `[1,1]
[2,2]
[3,3]
[4,4]
[5,5]
[6,6]`,
			expected: "[[[[5,0],[7,4]],[5,5]],[6,6]]",
		},
		{
			input: `[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]
[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]
[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]
[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]
[7,[5,[[3,8],[1,4]]]]
[[2,[2,2]],[8,[8,1]]]
[2,9]
[1,[[[9,3],9],[[9,0],[0,7]]]]
[[[5,[7,4]],7],1]
[[[[4,2],2],6],[8,7]]`,
			expected: "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
		},
	} {
		t.Run(tt.expected, func(t *testing.T) {
			answer := AddAll(challenge.FromLiteral(tt.input))

			require.Equal(t, tt.expected, answer.String())
		})
	}
}

func TestReduce(t *testing.T) {
	for _, tt := range []struct {
		line     string
		expected string
	}{
		{line: "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]", expected: "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"},
	} {
		t.Run(tt.line, func(t *testing.T) {
			sut, _ := Parse(tt.line)
			sut.reduce()

			require.Equal(t, tt.expected, sut.String())
		})
	}
}

func TestExplode(t *testing.T) {
	for _, tt := range []struct {
		line     string
		expected string
	}{
		{line: "[[[[[9,8],1],2],3],4]", expected: "[[[[0,9],2],3],4]"},
		{line: "[7,[6,[5,[4,[3,2]]]]]", expected: "[7,[6,[5,[7,0]]]]"},
		{line: "[[6,[5,[4,[3,2]]]],1]", expected: "[[6,[5,[7,0]]],3]"},
		{line: "[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]", expected: "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"},
		{line: "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]", expected: "[[3,[2,[8,0]]],[9,[5,[7,0]]]]"},
		{line: "[[[[4,0],[5,4]],[[7,0],[15,5]]],[10,[[11,9],[0,[11,8]]]]]", expected: "[[[[4,0],[5,4]],[[7,0],[15,5]]],[10,[[11,9],[11,0]]]]"},
	} {
		t.Run(tt.line, func(t *testing.T) {
			sut, _ := Parse(tt.line)

			assert.True(t, sut.explode())
			assert.Equal(t, tt.expected, sut.String())
		})
	}
}

func TestSplit(t *testing.T) {
	for _, tt := range []struct {
		v        int
		expected string
	}{
		{v: 10, expected: "[5,5]"},
		{v: 11, expected: "[5,6]"},
		{v: 12, expected: "[6,6]"},
	} {
		t.Run(tt.expected, func(t *testing.T) {
			sut := &snailfish{v: tt.v}

			assert.True(t, sut.split())
			assert.Equal(t, tt.expected, sut.String())
		})
	}
}

func TestMagnitude(t *testing.T) {
	for _, tt := range []struct {
		line     string
		expected int
	}{
		{line: "[[1,2],[[3,4],5]]", expected: 143},
		{line: "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", expected: 1384},
		{line: "[[[[1,1],[2,2]],[3,3]],[4,4]]", expected: 445},
		{line: "[[[[3,0],[5,3]],[4,4]],[5,5]]", expected: 791},
		{line: "[[[[5,0],[7,4]],[5,5]],[6,6]]", expected: 1137},
		{line: "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", expected: 3488},
	} {
		t.Run(tt.line, func(t *testing.T) {
			sut, _ := Parse(tt.line)

			require.Equal(t, tt.expected, sut.magnitude())
		})
	}
}
