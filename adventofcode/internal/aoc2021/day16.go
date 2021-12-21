package aoc2021

import (
	"bufio"

	"github.com/hierynomus/code-challenges/adventofcode/internal/aoc2021/bits"
	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day16(reader *bufio.Scanner) (string, string) {
	var part1, part2 int64

	inp := aoc.Read(reader)

	p, err := bits.NewPacketReader(inp).Next()
	if err != nil {
		panic(err)
	}

	part1 = AddVersions(p)
	part2 = p.Value()
	return aoc.Int64ToString(part1), aoc.Int64ToString(part2)
}

func AddVersions(p bits.Packet) int64 {
	var sum int64
	sum += int64(p.Version())

	switch p := p.(type) {
	case *bits.LiteralPacket:
	case *bits.OperatorPacket:
		for _, sp := range p.SubPackets {
			sum += AddVersions(sp)
		}
	}

	return sum
}
