package aoc2019

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

func TestDay06(t *testing.T) {
	d := day.TestDay(t, Day06)
	d.WithFile("../../input/aoc2019/day06.in", "117672", "277")
}
