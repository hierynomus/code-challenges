package aoc2020

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

func TestDay15_sample(t *testing.T) {
	inp := `0,3,6
`

	d := day.TestDay(t, Day15)
	d.WithInput(inp, "436", "175594")
}
