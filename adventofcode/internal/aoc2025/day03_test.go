package aoc2025

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

const A2025D03Sample = `987654321111111
811111111111119
234234234234278
818181911112111`

func TestDay03_Sample(t *testing.T) {
	d := day.TestDay(t, Day03)
	d.WithInput(A2025D03Sample, "357", "3121910778619")
}
