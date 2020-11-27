package aoc2019

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

func TestDay16_Real(t *testing.T) {
	d := day.TestDay(t, Day16)
	d.WithFile("../../input/aoc2019/day16.in", "90744714", "")
}
