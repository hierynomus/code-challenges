package aoc2019

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

func TestDay15_Real(t *testing.T) {
	d := day.TestDay(t, Day15)
	d.WithFile("../../input/aoc2019/day15.in", "258", "372")
}
