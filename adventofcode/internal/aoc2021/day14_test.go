package aoc2021

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

const TestInput = `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C
`

func TestDay14_sample(t *testing.T) {
	d := day.TestDay(t, Day14)
	d.WithInput(TestInput, "1588", "2188189693529")
}
