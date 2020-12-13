package aoc2020

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
	"github.com/stretchr/testify/assert"
)

func TestDay12_sample(t *testing.T) {
	inp := `F10
N3
F7
R90
F11
`

	d := day.TestDay(t, Day12)
	d.WithInput(inp, "25", "")
}

func TestTurn(t *testing.T) {
	assert.Equal(t, 'N', turn('S', 180, 1))
	assert.Equal(t, 'W', turn('S', 90, 1))
	assert.Equal(t, 'E', turn('S', 90, -1))
	assert.Equal(t, 'N', turn('S', 180, -1))
	assert.Equal(t, 'S', turn('N', 180, 1))
	assert.Equal(t, 'S', turn('N', 180, -1))
	assert.Equal(t, 'E', turn('W', 180, -1))
}
