package aoc2023

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

const A2023D11Sample = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....
`

func TestDay11_Sample(t *testing.T) {
	d := day.TestDay(t, Day11)
	d.WithInput(A2023D11Sample, "374", "82000210")
}
