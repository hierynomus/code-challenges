package aoc2023

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

const A2023D14Sample = `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....
`

func TestDay14_Sample(t *testing.T) {
	d := day.TestDay(t, Day14)
	d.WithInput(A2023D14Sample, "136", "64")
}
