package aoc2022

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

const A2022D17Sample = `>>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>
`

func TestDay17_Sample(t *testing.T) {
	d := day.TestDay(t, Day17)
	d.WithInput(A2022D17Sample, "3068", "")
}
