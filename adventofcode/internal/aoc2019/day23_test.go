package aoc2019

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

func TestDay23_Real(t *testing.T) {
	d := day.TestDay(t, Day23)
	d.WithFile("../../input/aoc2019/day23.in", "24106", "1140147758")
}
