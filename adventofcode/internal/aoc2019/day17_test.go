package aoc2019

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

func TestDay17_Real(t *testing.T) {
	d := day.TestDay(t, Day17)
	d.WithFile("../../input/aoc2019/day17.in", "3888", "927809")
}
