package aoc2021

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

const D17_Sample = `target area: x=20..30, y=-10..-5`

func TestDay17_Sample(t *testing.T) {
	d := day.TestDay(t, Day17)
	d.WithInput(D17_Sample, "45", "112")
}
