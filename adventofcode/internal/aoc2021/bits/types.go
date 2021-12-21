package bits

type PacketType int

const (
	SumPacketType PacketType = iota
	ProductPacketType
	MinimumPacketType
	MaximumPacketType
	LiteralPacketType
	GreaterThanPacketType
	LessThanPacketType
	EqualToPacketType
)

type Packet interface {
	Version() int
	Type() PacketType
	Value() int64
}

var _ Packet = (*LiteralPacket)(nil)
var _ Packet = (*OperatorPacket)(nil)

type LiteralPacket struct {
	version    int
	packetType PacketType
	value      int64
}

func (lp *LiteralPacket) Version() int {
	return lp.version
}

func (lp *LiteralPacket) Type() PacketType {
	return lp.packetType
}

func (lp *LiteralPacket) Value() int64 {
	return lp.value
}

type OperatorPacket struct {
	version    int
	packetType PacketType
	SubPackets []Packet
}

func (op *OperatorPacket) Version() int {
	return op.version
}

func (op *OperatorPacket) Type() PacketType {
	return op.packetType
}

func (op *OperatorPacket) Value() int64 {
	switch op.packetType {
	case SumPacketType:
		return sum(op.SubPackets)
	case ProductPacketType:
		return product(op.SubPackets)
	case MinimumPacketType:
		return min(op.SubPackets)
	case MaximumPacketType:
		return max(op.SubPackets)
	case GreaterThanPacketType:
		if op.SubPacket(0).Value() > op.SubPacket(1).Value() {
			return 1
		} else {
			return 0
		}
	case LessThanPacketType:
		if op.SubPacket(0).Value() < op.SubPacket(1).Value() {
			return 1
		} else {
			return 0
		}
	case EqualToPacketType:
		if op.SubPacket(0).Value() == op.SubPacket(1).Value() {
			return 1
		} else {
			return 0
		}
	default:
		panic("Unknown operator")
	}
}

func (op *OperatorPacket) NrPackets() int {
	return len(op.SubPackets)
}

func (op *OperatorPacket) SubPacket(i int) Packet {
	return op.SubPackets[i]
}

func sum(packets []Packet) int64 {
	s := int64(0)
	for _, p := range packets {
		s += p.Value()
	}
	return s
}

func product(packets []Packet) int64 {
	s := int64(1)
	for _, p := range packets {
		s *= p.Value()
	}
	return s
}

func min(packets []Packet) int64 {
	s := int64(1 << 62)
	for _, p := range packets {
		v := p.Value()
		if v < s {
			s = v
		}
	}
	return s
}

func max(packets []Packet) int64 {
	s := int64(0)
	for _, p := range packets {
		v := p.Value()
		if v > s {
			s = v
		}
	}
	return s
}
