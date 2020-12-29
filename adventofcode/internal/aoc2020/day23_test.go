package aoc2020

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
	"github.com/stretchr/testify/assert"
)

func TestDay23_sample(t *testing.T) {
	inp := "389125467"
	m, c := parseCups(inp)
	playCrabCups(m, c, 10)
	arr := []int{}
	cup := m[1]
	for i := 0; i < 8; i++ {
		arr = append(arr, cup.Next.Nr)
		cup = cup.Next
	}
	assert.Equal(t, aoc.IntArrayAsString(arr, ""), "92658374")
}
