package aoc2019

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

func TestDay12_1(t *testing.T) {
	input := `<x=-8, y=-10, z=0>
<x=5, y=5, z=10>
<x=2, y=-7, z=3>
<x=9, y=-8, z=-3>
`
	d := day.TestDay(t, Day12)
	// d.WithInput(input, "14645", "4686774924")
	d.WithInput(input, "14645", "")
}

func TestDay12_2(t *testing.T) {
	input := `<x=-1, y=0, z=2>
<x=2, y=-10, z=-7>
<x=4, y=-8, z=8>
<x=3, y=5, z=-1>
`
	d := day.TestDay(t, Day12)
	// d.WithInput(input, "183", "2772")
	d.WithInput(input, "183", "")
}
