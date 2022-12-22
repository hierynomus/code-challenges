package aoc2022

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

const A2022D20Sample = `1
2
-3
3
-2
0
4
`

func TestDay20_Sample(t *testing.T) {
	d := day.TestDay(t, Day20)
	d.WithInput(A2022D20Sample, "3", "")
}
