package aoc2019

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

func TestDay18_Real(t *testing.T) {
	d := day.TestDay(t, Day18)
	d.WithFile("../../input/aoc2019/day18.in", "90744714", "")
}
