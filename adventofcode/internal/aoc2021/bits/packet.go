package bits

import (
	"errors"
	"strconv"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type PacketReader struct {
	r *Reader
}

func NewPacketReader(inp string) *PacketReader {
	return &PacketReader{NewTransmissionReader(inp)}
}

func (pr *PacketReader) Next() (Packet, error) {
	if pr.r.Done() {
		return nil, errors.New("no more packets available")
	}

	version, err := pr.readVersion()
	if err != nil {
		return nil, err
	}

	t, err := pr.readType()
	if err != nil {
		return nil, err
	}

	switch t {
	case LiteralPacketType:
		return pr.readLiteralPacket(version, t)
	default:
		return pr.readOperatorPacket(version, t)
	}
}

func (pr *PacketReader) readLiteralPacket(version int, packetType PacketType) (Packet, error) {
	p := &LiteralPacket{version: version, packetType: packetType}
	nr := strings.Builder{}
	for {
		b, err := pr.r.ReadBit()
		if err != nil {
			return nil, err
		}

		v, err := pr.r.ReadBits(4)
		if err != nil {
			return nil, err
		}
		nr.WriteString(string(v))
		if b == '0' {
			break
		}
	}

	val, err := strconv.ParseInt(nr.String(), 2, 64)
	if err != nil {
		return nil, err
	}

	p.value = val
	return p, nil
}

func (pr *PacketReader) readOperatorPacket(version int, packetType PacketType) (Packet, error) {
	p := &OperatorPacket{version: version, packetType: packetType}

	lengthType, err := pr.r.ReadBit()
	if err != nil {
		return nil, err
	}

	subPackets := []Packet{}
	if lengthType == '0' {
		length, err := pr.readSubPacketLength()
		if err != nil {
			return nil, err
		}

		curPos := pr.r.pos
		for curPos+length > pr.r.pos {
			subPacket, err := pr.Next()
			if err != nil {
				return nil, err
			}
			subPackets = append(subPackets, subPacket)
		}
	} else if lengthType == '1' {
		nr, err := pr.readNrSubPackets()
		if err != nil {
			return nil, err
		}

		for i := 0; i < nr; i++ {
			p, err := pr.Next()
			if err != nil {
				return nil, err
			}

			subPackets = append(subPackets, p)
		}
	}

	p.SubPackets = subPackets
	return p, nil
}

func (pr *PacketReader) readVersion() (int, error) {
	v, err := pr.r.ReadBits(3)
	if err != nil {
		return -1, err
	}

	i, err := strconv.ParseInt(string(v), 2, 8)
	if err != nil {
		return -1, err
	}

	return int(i), nil
}

func (pr *PacketReader) readType() (PacketType, error) {
	v, err := pr.r.ReadBits(3)
	if err != nil {
		return -1, err
	}

	return PacketType(aoc.BinaryToInt8(string(v))), nil
}

func (pr *PacketReader) readSubPacketLength() (int, error) {
	v, err := pr.r.ReadBits(15)
	if err != nil {
		return -1, err
	}

	return aoc.BinaryToInt16(string(v)), nil
}

func (pr *PacketReader) readNrSubPackets() (int, error) {
	v, err := pr.r.ReadBits(11)
	if err != nil {
		return -1, err
	}

	return aoc.BinaryToInt16(string(v)), nil
}
