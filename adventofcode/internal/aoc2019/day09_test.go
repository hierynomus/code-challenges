package aoc2019

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

func TestDay09_Real(t *testing.T) {
	d := day.TestDay(t, Day09)
	d.WithFile("../../input/aoc2019/day09.in", "2427443564", "87221")
}
