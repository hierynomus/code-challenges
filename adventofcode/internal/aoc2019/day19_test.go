package aoc2019

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

func TestDay19_Real(t *testing.T) {
	d := day.TestDay(t, Day19)
	d.WithFile("../../input/aoc2019/day19.in", "234", "9290812")
}
