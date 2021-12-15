package aoc2021

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

func TestDay15_sample(t *testing.T) {
	inp := `1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581`

	d := day.TestDay(t, Day15)
	d.WithInput(inp, "40", "315")
}
