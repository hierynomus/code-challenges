package aoc2023

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

const A2023D06Sample = `Time:      7  15   30
Distance:  9  40  200`

func TestDay06_Sample(t *testing.T) {
	d := day.TestDay(t, Day06)
	d.WithInput(A2023D06Sample, "288", "71503")
}
