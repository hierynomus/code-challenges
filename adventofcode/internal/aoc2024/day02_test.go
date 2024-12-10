package aoc2024

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

const A2024D02Sample = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func TestDay02_Sample(t *testing.T) {
	d := day.TestDay(t, Day02)
	d.WithInput(A2024D02Sample, "2", "4")
}
