package aoc2019

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

func TestDay13_Real(t *testing.T) {
	d := day.TestDay(t, Day13)
	d.WithFile("../../input/aoc2019/day13.in", "361", "17590")
}
