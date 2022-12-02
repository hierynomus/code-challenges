package bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPackets(t *testing.T) {
	cases := []struct {
		inp   string
		value int64
	}{
		{"C200B40A82", 3},
		{"04005AC33890", 54},
		{"880086C3E88112", 7},
		{"CE00C43D881120", 9},
		{"D8005AC2A8F0", 1},
		{"F600BC2D8F", 0},
		{"9C005AC2F8F0", 0},
		{"9C0141080250320F1802104A08", 1},
	}
	for _, c := range cases {
		pr := NewPacketReader(c.inp)
		pkt, err := pr.Next()
		assert.NoError(t, err)
		assert.Equal(t, c.value, pkt.Value())
	}
}

func TestReadOperatorPacketNrPackets(t *testing.T) {
	p := "EE00D40C823060"
	pr := NewPacketReader(p)
	pkt, err := pr.Next()
	assert.NoError(t, err)
	assert.Equal(t, 7, pkt.Version())
	assert.Equal(t, PacketType(3), pkt.Type())
	assert.Equal(t, 3, pkt.(*OperatorPacket).NrPackets())
	assert.Equal(t, int64(1), pkt.(*OperatorPacket).SubPacket(0).Value())
	assert.Equal(t, int64(2), pkt.(*OperatorPacket).SubPacket(1).Value())
	assert.Equal(t, int64(3), pkt.(*OperatorPacket).SubPacket(2).Value())
}

func TestReadOperatorPacketLength(t *testing.T) {
	p := "38006F45291200"
	pr := NewPacketReader(p)
	pkt, err := pr.Next()
	assert.NoError(t, err)
	assert.Equal(t, 1, pkt.Version())
	assert.Equal(t, PacketType(6), pkt.Type())
	assert.Equal(t, 2, pkt.(*OperatorPacket).NrPackets())
	assert.Equal(t, int64(10), pkt.(*OperatorPacket).SubPacket(0).Value())
	assert.Equal(t, int64(20), pkt.(*OperatorPacket).SubPacket(1).Value())
}

func TestReadLiteralPacket(t *testing.T) {
	p := "D2FE28"
	pr := NewPacketReader(p)
	pkt, err := pr.Next()
	assert.NoError(t, err)
	assert.Equal(t, pkt.Version(), 6)
	assert.Equal(t, pkt.Type(), LiteralPacketType)
	assert.Equal(t, pkt.Value(), int64(2021))
}
