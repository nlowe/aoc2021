package day12

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/nlowe/aoc2021/challenge"
)

const (
	smallExample = `start-A
start-b
A-c
A-b
b-d
A-end
b-end`

	mediumExample = `dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`

	largeExample = `fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW`
)

func TestA(t *testing.T) {
	for _, tt := range []struct {
		name     string
		edges    string
		expected int
	}{
		{name: "small", edges: smallExample, expected: 10},
		{name: "medium", edges: mediumExample, expected: 19},
		{name: "large", edges: largeExample, expected: 226},
	} {
		t.Run(tt.name, func(t *testing.T) {
			input := challenge.FromLiteral(tt.edges)

			result := partA(input)

			require.Equal(t, tt.expected, result)
		})
	}
}
