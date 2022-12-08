package aoc2022

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

const A2022D08Sample = `30373
25512
65332
33549
35390`

func TestDay08_Sample(t *testing.T) {
	d := day.TestDay(t, Day08)
	d.WithInput(A2022D08Sample, "21", "8")
}
