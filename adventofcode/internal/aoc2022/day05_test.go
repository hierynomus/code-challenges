package aoc2022

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

const A2022D05Sample = `    [D]
[N] [C]
[Z] [M] [P]
 1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

func TestDay05_Sample(t *testing.T) {
	d := day.TestDay(t, Day05)
	d.WithInput(A2022D05Sample, "CMZ", "MCD")
}
