package aoc2024

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

const A2024D10Sample = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

func TestDay10_Sample(t *testing.T) {
	d := day.TestDay(t, Day10)
	d.WithInput(A2024D10Sample, "36", "81")
}
