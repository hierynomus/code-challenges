package aoc2018

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

func TestDay05_sample(t *testing.T) {
	inp := `dabAcCaCBAcCcaDA`

	d := day.TestDay(t, Day05)
	d.WithInput(inp, "10", "4")
}
