package aoc2019

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

func TestDay21_Real(t *testing.T) {
	d := day.TestDay(t, Day21)
	d.WithFile("../../input/aoc2019/day21.in", "19357335", "1140147758")
}
