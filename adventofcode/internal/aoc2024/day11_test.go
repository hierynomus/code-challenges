package aoc2024

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
	"github.com/stretchr/testify/assert"
)

const A2024D11Sample = `125 17`

func TestDay11_Sample(t *testing.T) {
	d := day.TestDay(t, Day11)
	d.WithInput(A2024D11Sample, "55312", "81")
}

func TestDay11_Blink(t *testing.T) {
	stones := map[int]int{125: 1, 17: 1}
	stones = blink(stones)
	assert.Equal(t, map[int]int{253000: 1, 1: 1, 7: 1}, stones)
}
