package aoc2023

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

const A2023D08Sample = `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`

func TestDay08_Sample(t *testing.T) {
	d := day.TestDay(t, Day08)
	d.WithInput(A2023D08Sample, "2", "0")
}

func TestDay08_Sample2(t *testing.T) {
	d := day.TestDay(t, Day08)
	d.WithInput(`LLR

	AAA = (BBB, BBB)
	BBB = (AAA, ZZZ)
	ZZZ = (ZZZ, ZZZ)`, "6", "0")
}
