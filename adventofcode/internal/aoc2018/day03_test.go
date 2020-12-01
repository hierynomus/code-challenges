package aoc2018

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

func TestDay03_sample(t *testing.T) {
	inp := `#1 @ 1,3: 4x4
#2 @ 3,1: 4x4
#3 @ 5,5: 2x2`

	d := day.TestDay(t, Day03)
	d.WithInput(inp, "4", "3")
}
