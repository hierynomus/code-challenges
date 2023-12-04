package aoc2022

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

const A2022D23Sample = `....#..
..###.#
#...#.#
.#...##
#.###..
##.#.##
.#..#..
`

func TestDay23_Sample(t *testing.T) {
	d := day.TestDay(t, Day23)
	d.WithInput(A2022D23Sample, "110", "20")
}
