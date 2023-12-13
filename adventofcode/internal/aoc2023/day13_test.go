package aoc2023

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

const A2023D13Sample = `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#
`

func TestDay13_Sample(t *testing.T) {
	d := day.TestDay(t, Day13)
	d.WithInput(A2023D13Sample, "405", "400")
}
