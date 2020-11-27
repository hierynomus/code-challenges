package aoc2019

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

func TestDay24_Real(t *testing.T) {
	d := day.TestDay(t, Day24)
	d.WithFile("../../input/aoc2019/day24.in", "18370591", "")
}
