package aoc2021

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

func TestDay07_sample(t *testing.T) {
	inp := `16,1,2,0,4,2,7,1,2,14`

	d := day.TestDay(t, Day07)
	d.WithInput(inp, "37", "168")
}
