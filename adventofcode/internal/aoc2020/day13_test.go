package aoc2020

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

func TestDay13_sample(t *testing.T) {
	inp := `939
7,13,x,x,59,x,31,19
`

	d := day.TestDay(t, Day13)
	d.WithInput(inp, "295", "1068781")
}
func TestDay13_sample2(t *testing.T) {
	inp := `939
17,x,13,19
`

	d := day.TestDay(t, Day13)
	d.WithInput(inp, "130", "3417")
}
