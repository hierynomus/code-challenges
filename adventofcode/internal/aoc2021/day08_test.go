package aoc2021

import (
	"fmt"
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
	"github.com/stretchr/testify/assert"
)

func TestDay08_sample(t *testing.T) {
	inp := `acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf`

	d := day.TestDay(t, Day08)
	d.WithInput(inp, "0", "5353")
}

func TestDemangle(t *testing.T) {
	ss := Demangle("acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab")
	fmt.Printf("%v", ss)
	println(ss.RenderNumber("cdfeb"))
	assert.Equal(t, 5, ss.AsNumber("cdfeb"))
	println(ss.RenderNumber("fcadb"))
	assert.Equal(t, 3, ss.AsNumber("fcadb"))
	println(ss.RenderNumber("cdfeb"))
	println(ss.RenderNumber("cdbaf"))
	assert.Equal(t, 3, ss.AsNumber("cdbaf"))

}
