package aoc2023

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
	"github.com/stretchr/testify/assert"
)

const A2023D07Sample = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

func TestDay07_Sample(t *testing.T) {
	d := day.TestDay(t, Day07)
	d.WithInput(A2023D07Sample, "6440", "5905")
}

func TestCamelCardHandType(t *testing.T) {
	assert.Equal(t, ThreeOfAKind, parseHandType("6JQ9J", true))
	assert.Equal(t, OnePair, parseHandType("6JQ9J", false))
}
