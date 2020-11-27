package aoc2019

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

func TestDay07_Real(t *testing.T) {
	d := day.TestDay(t, Day07)
	d.WithFile("../../input/aoc2019/day07.in", "929800", "15432220")
}
