package intcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func TestShouldFunctionParameterMode(t *testing.T) {
	input := aoc.AsIntArray("1002,4,3,4,33")
	icm := NewIntCodeMachine(input)
	i := icm.Run()
	assert.Equal(t, i, 1002)
}

func TestShouldJumpCorrectlyPositionMode(t *testing.T) {
	input := aoc.AsIntArray("3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9")
	icm := NewIntCodeMachine(input)

	go icm.Run()
	icm.Input.Write(0)
	assert.Equal(t, icm.Output.Read(), 0)
	icm.Reset()

	go icm.Run()
	icm.Input.Write(2)
	assert.Equal(t, icm.Output.Read(), 1)
}

func TestShouldJumpCorrectlyImmediateMode(t *testing.T) {
	input := aoc.AsIntArray("3,3,1105,-1,9,1101,0,0,12,4,12,99,1")
	icm := NewIntCodeMachine(input)

	go icm.Run()
	icm.Input.Write(0)
	assert.Equal(t, icm.Output.Read(), 0)
	icm.Reset()

	go icm.Run()
	icm.Input.Write(2)
	assert.Equal(t, icm.Output.Read(), 1)
}

func TestIntCodeDay05_1(t *testing.T) {
	input := aoc.AsIntArray("3,9,8,9,10,9,4,9,99,-1,8")
	icm := NewIntCodeMachine(input)

	go icm.Run()
	icm.Input.Write(8)
	assert.Equal(t, icm.Output.Read(), 1)
}

func TestIntCodeDay09_1(t *testing.T) {
	input := aoc.AsIntArray("109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99")
	icm := NewIntCodeMachine(input)

	go icm.Run()

	out := []int{}

loop:
	for {
		select {
		case <-icm.ClosedCh:
			break loop
		default:
			out = append(out, icm.Output.Read())
		}
	}

	assert.Equal(t, input, out)
}

func TestIntCodeDay09_2(t *testing.T) {
	input := aoc.AsIntArray("1102,34915192,34915192,7,4,7,99,0")
	icm := NewIntCodeMachine(input)

	go icm.Run()
	assert.Equal(t, icm.Output.Read(), 1219070632396864)
}

func TestIntCodeDay09_3(t *testing.T) {
	input := aoc.AsIntArray("104,1125899906842624,99")
	icm := NewIntCodeMachine(input)

	go icm.Run()
	assert.Equal(t, icm.Output.Read(), 1125899906842624)
}
