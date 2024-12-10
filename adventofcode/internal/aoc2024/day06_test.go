package aoc2024

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

const A2024D06Sample = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

func TestDay06_Sample(t *testing.T) {
	d := day.TestDay(t, Day06)
	d.WithInput(A2024D06Sample, "41", "6")
}
