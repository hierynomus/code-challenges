package aoc2023

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

const A2023D10Sample = `-L|F7
7S-7|
L|7||
-L-J|
L|-JF
`

func TestDay10_Sample(t *testing.T) {
	d := day.TestDay(t, Day10)
	d.WithInput(A2023D10Sample, "4", "0")
}

func TestDay10_Sample2(t *testing.T) {
	d := day.TestDay(t, Day10)
	d.WithInput(`7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ
`, "8", "0")
}
