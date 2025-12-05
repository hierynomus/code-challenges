package aoc2025

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

const A2025D05Sample = `3-5
10-14
16-20
12-18

1
5
8
11
17
32
`

func TestDay05_Sample(t *testing.T) {
	d := day.TestDay(t, Day05)
	d.WithInput(A2025D05Sample, "3", "14")
}
