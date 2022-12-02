package bits

import (
	"errors"
)

type Transmission []rune

type Reader struct {
	transmission Transmission
	pos          int
}

var HexBin map[rune]string = map[rune]string{
	'0': "0000",
	'1': "0001",
	'2': "0010",
	'3': "0011",
	'4': "0100",
	'5': "0101",
	'6': "0110",
	'7': "0111",
	'8': "1000",
	'9': "1001",
	'A': "1010",
	'B': "1011",
	'C': "1100",
	'D': "1101",
	'E': "1110",
	'F': "1111",
}

func NewTransmissionReader(inp string) *Reader {
	return &Reader{toBinaryRepresentation(inp), 0}
}

func toBinaryRepresentation(s string) []rune {
	var r []rune
	for _, c := range s {
		r = append(r, []rune(HexBin[c])...)
	}
	return r
}

func (r *Reader) ReadBit() (rune, error) {
	if r.Done() {
		return rune(0), errors.New("bits are empty")
	}
	c := r.transmission[r.pos]
	r.pos++
	return c, nil
}

func (r *Reader) ReadBits(nr int) ([]rune, error) {
	if r.pos+nr > len(r.transmission) {
		return nil, errors.New("not enough bits")
	}
	bits := r.transmission[r.pos : r.pos+nr]
	r.pos += nr
	return bits, nil
}

func (r *Reader) Done() bool {
	return r.pos >= len(r.transmission)
}
