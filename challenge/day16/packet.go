package day16

import (
	"encoding/hex"
	"fmt"

	"github.com/nlowe/aoc2021/challenge"
)

const (
	packetTypeSum     = 0
	packetTypeMul     = 1
	packetTypeMin     = 2
	packetTypeMax     = 3
	packetTypeLiteral = 4
	packetTypeGreater = 5
	packetTypeLess    = 6
	packetTypeEqual   = 7

	literalEndOfGroupsMask = 0b10000
	literalGroupMask       = 0b1111
)

type packet interface {
	Version() byte
	TypeID() byte
	VersionSum() int
	Eval() int
}

type packetHeader struct {
	version byte
	typeID  byte
}

func (p packetHeader) Version() byte {
	return p.version
}

func (p packetHeader) TypeID() byte {
	return p.typeID
}

type literalPacket struct {
	packetHeader

	v int
}

func (l literalPacket) VersionSum() int {
	return int(l.version)
}

func (l literalPacket) Eval() int {
	return l.v
}

type operatorPacket struct {
	packetHeader

	lengthType bool

	subPackets []packet
}

func (o operatorPacket) VersionSum() int {
	sum := int(o.version)

	for _, sub := range o.subPackets {
		sum += sub.VersionSum()
	}

	return sum
}

func (o operatorPacket) Eval() (result int) {
	switch o.typeID {
	case packetTypeSum:
		for _, sub := range o.subPackets {
			result += sub.Eval()
		}
	case packetTypeMul:
		result = 1
		for _, sub := range o.subPackets {
			result *= sub.Eval()
		}
	case packetTypeMin:
		result = o.subPackets[0].Eval()

		for i, sub := range o.subPackets {
			if i == 0 {
				continue
			}

			v := sub.Eval()
			if v < result {
				result = v
			}
		}
	case packetTypeMax:
		for _, sub := range o.subPackets {
			v := sub.Eval()

			if v > result {
				result = v
			}
		}
	case packetTypeGreater:
		if o.subPackets[0].Eval() > o.subPackets[1].Eval() {
			result = 1
		}
	case packetTypeLess:
		if o.subPackets[0].Eval() < o.subPackets[1].Eval() {
			result = 1
		}
	case packetTypeEqual:
		if o.subPackets[0].Eval() == o.subPackets[1].Eval() {
			result = 1
		}
	default:
		panic(fmt.Errorf("unknown packet type id %d", o.typeID))
	}

	return
}

// bitstream is a wrapper over []byte that allows for reading up to 8 bits
// at a time. The bitstream panics if you try to read more bits than it contains.
type bitstream struct {
	// b contains the bytes to read from
	b []byte

	// off is the offset into b
	off int
	// sub is the bit offset into the byte at b[off] where 0 is the most significant bit
	sub int
}

func NewBitstream(challenge *challenge.Input) *bitstream {
	bytes, err := hex.DecodeString(<-challenge.Lines())
	if err != nil {
		panic(err)
	}

	return &bitstream{b: bytes}
}

func (b *bitstream) read(bits int) (result int) {
	if bits >= 64 {
		panic("read too large")
	}

	for i := 0; i < bits; i++ {
		if b.off >= len(b.b) {
			panic(fmt.Errorf("out of bounds read at [%d]+%d for %d bytes", b.off, b.sub, len(b.b)))
		}

		result <<= 1
		result |= int((b.b[b.off] >> (7 - b.sub)) & 0b1)

		b.sub++
		if b.sub == 8 {
			b.sub = 0
			b.off++
		}
	}

	return
}

func (b *bitstream) parse() (packet, int) {
	version := byte(b.read(3))
	typeID := byte(b.read(3))

	h := packetHeader{version, typeID}

	switch typeID {
	case packetTypeLiteral:
		p, c := b.parseLiteral(h)
		return p, c + 6
	default:
		p, c := b.parseOperator(h)
		return p, c + 6
	}
}

func (b *bitstream) parseLiteral(h packetHeader) (literalPacket, int) {
	read := 0

	var v int
	for {
		group := b.read(5)
		read += 5

		more := group&literalEndOfGroupsMask == literalEndOfGroupsMask
		subNumber := group & literalGroupMask

		v <<= 4
		v |= subNumber

		// last group?
		if !more {
			break
		}
	}

	return literalPacket{
		packetHeader: h,
		v:            v,
	}, read
}

func (b *bitstream) parseOperator(h packetHeader) (operatorPacket, int) {
	lengthType := b.read(1) == 0b1
	read := 1

	result := operatorPacket{
		packetHeader: h,
		lengthType:   lengthType,
		subPackets:   nil,
	}

	if lengthType {
		// next 11 bits are the number of sub-packets
		subPackets := b.read(11)
		read += 11

		for i := 0; i < subPackets; i++ {
			sub, consumed := b.parse()
			read += consumed

			result.subPackets = append(result.subPackets, sub)
		}
	} else {
		// next 15 bits are the total length in bits of sub-packets
		subPacketBits := b.read(15)
		read += 15

		for subPacketBits > 0 {
			sub, consumed := b.parse()
			read += consumed
			subPacketBits -= consumed

			result.subPackets = append(result.subPackets, sub)
		}
	}

	return result, read
}
