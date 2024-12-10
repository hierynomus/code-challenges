package aoc2024

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

const A2024D09Sample = `2333133121414131402`

func TestDay09_Sample(t *testing.T) {
	d := day.TestDay(t, Day09)
	d.WithInput(A2024D09Sample, "1928", "2858")
}
