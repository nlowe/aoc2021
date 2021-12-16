package day16

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/nlowe/aoc2021/challenge"
)

func TestA(t *testing.T) {
	for _, tt := range []struct {
		input    string
		expected int
	}{
		{input: "8A004A801A8002F478", expected: 16},
		{input: "620080001611562C8802118E34", expected: 12},
		{input: "C0015000016115A2E0802F182340", expected: 23},
		{input: "A0016C880162017C3686B18A3D4780", expected: 31},
	} {
		t.Run(tt.input, func(t *testing.T) {
			input := challenge.FromLiteral(tt.input)

			result := partA(input)

			require.Equal(t, tt.expected, result)
		})
	}
}

func TestParse(t *testing.T) {
	t.Run("basic literal", func(t *testing.T) {
		bytes, err := hex.DecodeString("D2FE28")
		require.NoError(t, err)

		sut := bitstream{b: bytes}

		p, _ := sut.parse()
		require.IsType(t, literalPacket{}, p)

		l := p.(literalPacket)

		assert.Equal(t, byte(6), l.version)
		assert.Equal(t, 2021, l.v)
	})

	t.Run("bit operator", func(t *testing.T) {
		bytes, err := hex.DecodeString("38006F45291200")
		require.NoError(t, err)

		sut := bitstream{b: bytes}

		p, _ := sut.parse()
		require.IsType(t, operatorPacket{}, p)

		o := p.(operatorPacket)

		assert.Equal(t, byte(1), o.version)
		assert.Equal(t, byte(6), o.typeID)

		require.Len(t, o.subPackets, 2)

		lp0 := o.subPackets[0]
		require.IsType(t, literalPacket{}, lp0)
		l0 := lp0.(literalPacket)
		assert.Equal(t, byte(6), l0.version)
		assert.Equal(t, 10, l0.v)

		lp1 := o.subPackets[1]
		require.IsType(t, literalPacket{}, lp1)
		l1 := lp1.(literalPacket)
		assert.Equal(t, byte(2), l1.version)
		assert.Equal(t, 20, l1.v)
	})

	t.Run("count operator", func(t *testing.T) {
		bytes, err := hex.DecodeString("EE00D40C823060")
		require.NoError(t, err)

		sut := bitstream{b: bytes}

		p, _ := sut.parse()
		require.IsType(t, operatorPacket{}, p)

		o := p.(operatorPacket)

		assert.Equal(t, byte(7), o.version)
		assert.Equal(t, byte(3), o.typeID)

		require.Len(t, o.subPackets, 3)

		lp0 := o.subPackets[0]
		require.IsType(t, literalPacket{}, lp0)
		l0 := lp0.(literalPacket)
		assert.Equal(t, byte(2), l0.version)
		assert.Equal(t, 1, l0.v)

		lp1 := o.subPackets[1]
		require.IsType(t, literalPacket{}, lp1)
		l1 := lp1.(literalPacket)
		assert.Equal(t, byte(4), l1.version)
		assert.Equal(t, 2, l1.v)

		lp2 := o.subPackets[2]
		require.IsType(t, literalPacket{}, lp2)
		l2 := lp2.(literalPacket)
		assert.Equal(t, byte(1), l2.version)
		assert.Equal(t, 3, l2.v)
	})
}
