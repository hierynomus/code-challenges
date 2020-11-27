package aoc2019

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

func TestDay22_Real(t *testing.T) {
	d := day.TestDay(t, Day22)
	d.WithFile("../../input/aoc2019/day22.in", "7096", "")
}
