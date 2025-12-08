package aoc2025

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

const A2025D06Sample = "123 328  51 64 \n 45 64  387 23 \n  6 98  215 314\n*   +   *   +  \n"

func TestDay06_Sample(t *testing.T) {
	d := day.TestDay(t, Day06)
	d.WithInput(A2025D06Sample, "4277556", "3263827")
}
