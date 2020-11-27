package aoc2019

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

func TestDay05(t *testing.T) {
	d := day.TestDay(t, Day05)
	d.WithFile("../../input/aoc2019/day05.in", "15426686", "11430197")
}
