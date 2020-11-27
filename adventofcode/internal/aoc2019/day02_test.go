package aoc2019

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

func TestDay02(t *testing.T) {
	d := day.TestDay(t, Day02)
	d.WithFile("../../input/aoc2019/day02.in", "4138687", "6635")
}
