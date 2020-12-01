package aoc2019

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func TestHistogram(t *testing.T) {
	i := []int{1, 1, 2, 3, 4, 6, 8, 6, 3}
	h := aoc.MakeIntHistogram(i)
	assert.Equal(t, h[1], 2)
	assert.Equal(t, h[6], 2)
}
