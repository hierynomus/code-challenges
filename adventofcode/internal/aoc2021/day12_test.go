package aoc2021

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

func TestDay12_sample(t *testing.T) {
	inp := `start-A
start-b
A-c
A-b
b-d
A-end
b-end`

	d := day.TestDay(t, Day12)
	d.WithInput(inp, "10", "36")
}
