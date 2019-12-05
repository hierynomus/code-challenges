package intcode

import "testing"

import "github.com/hierynomus/aoc2019/aoc"

import "gotest.tools/v3/assert"

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
	icm.IO.Input <- 0
	assert.Equal(t, <-icm.IO.Output, 0)
	icm.Reset()
	go icm.Run()
	icm.IO.Input <- 2
	assert.Equal(t, <-icm.IO.Output, 1)
}

func TestShouldJumpCorrectlyImmediateMode(t *testing.T) {
	input := aoc.AsIntArray("3,3,1105,-1,9,1101,0,0,12,4,12,99,1")
	icm := NewIntCodeMachine(input)
	go icm.Run()
	icm.IO.Input <- 0
	assert.Equal(t, <-icm.IO.Output, 0)
	icm.Reset()
	go icm.Run()
	icm.IO.Input <- 2
	assert.Equal(t, <-icm.IO.Output, 1)
}
