package aoc2020

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
	"github.com/stretchr/testify/assert"
)

func TestDay23_sample(t *testing.T) {
	inp := []int{3, 8, 9, 1, 2, 5, 4, 6, 7}
	out := playCrabCups(inp, 10)
	assert.Equal(t, aoc.IntArrayAsString(out, ""), "92658374")
}
