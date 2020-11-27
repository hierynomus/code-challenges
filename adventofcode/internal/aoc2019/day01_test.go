package aoc2019

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

func TestDay01(t *testing.T) {
	d := day.TestDay(t, Day01)
	d.WithFile("../../input/aoc2019/day01.in", "3388015", "5079140")
}
