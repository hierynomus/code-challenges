package aoc2021

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

func TestDay21_sample(t *testing.T) {
	inp := `Player 1 starting position: 4
Player 2 starting position: 8`

	d := day.TestDay(t, Day21)
	d.WithInput(inp, "37", "168")
}
