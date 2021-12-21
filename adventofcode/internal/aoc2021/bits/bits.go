package bits

type Machine struct {
	r *Reader
}

func NewMachine(inp string) *Machine {
	return &Machine{NewTransmissionReader(inp)}
}

// func (m *Machine) Run() {
// 	for !m.r.Done() {
// 		c := m.r.Read()
// 		if c == '1' {
// 			packets = append(packets, LiteralPacket{})
// 		} else {
// 			packets = append(packets, OperatorPacket{})
// 		}
// 	}
// }

// func ParsePackets(reader *bufio.Scanner) []Packet {
// 	var packets []Packet
// 	for reader.Scan() {
// 		l := reader.Text()
// 	}
// 	return packets
// }
