package aoc2025

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

const A2025D04Sample = `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.
`

func TestDay04_Sample(t *testing.T) {
	d := day.TestDay(t, Day04)
	d.WithInput(A2025D04Sample, "13", "43")
}
