package aoc2023

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

const A2023D03Sample = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

func TestDay03_Sample(t *testing.T) {
	d := day.TestDay(t, Day03)
	d.WithInput(A2023D03Sample, "4361", "467835")
}

func TestDay03_IdenticalParts(t *testing.T) {
	d := day.TestDay(t, Day03)
	d.WithInput(`....................
..-52..52-..52..52..
..................-.`, "156", "0")
}
