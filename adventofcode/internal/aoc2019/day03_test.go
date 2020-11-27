package aoc2019

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

func TestDay03(t *testing.T) {
	d := day.TestDay(t, Day03)
	d.WithFile("../../input/aoc2019/day03.in", "446", "9006")
}
