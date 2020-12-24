package aoc2020

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

func TestDay22_sample(t *testing.T) {
	inp := `Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10
`

	d := day.TestDay(t, Day22)
	d.WithInput(inp, "306", "291")
}
